"""
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""

from typing import Any, Callable, Dict, List, Optional, Type

from magma.common.service import MagmaService
from magma.enodebd.data_models import transform_for_enb, transform_for_magma
from magma.enodebd.data_models.data_model import DataModel, TrParam
from magma.enodebd.data_models.data_model_parameters import (
    ParameterName,
    TrParameterType,
)
from magma.enodebd.device_config.configuration_init import build_desired_config
from magma.enodebd.device_config.enodeb_config_postprocessor import (
    EnodebConfigurationPostProcessor,
)
from magma.enodebd.device_config.enodeb_configuration import EnodebConfiguration
from magma.enodebd.devices.device_utils import EnodebDeviceName
from magma.enodebd.logger import EnodebdLogger as logger
from magma.enodebd.state_machines.acs_state_utils import (
    get_all_objects_to_add,
    get_all_objects_to_delete,
    get_all_param_values_to_set,
    get_params_to_get,
    parse_get_parameter_values_response,
)
from magma.enodebd.state_machines.enb_acs_impl import BasicEnodebAcsStateMachine
from magma.enodebd.state_machines.enb_acs_states import (
    AcsMsgAndTransition,
    AcsReadMsgResult,
    AddObjectsState,
    BaicellsSendRebootState,
    CheckOptionalParamsState,
    DeleteObjectsState,
    EndSessionState,
    EnodebAcsState,
    ErrorState,
    GetObjectParametersState,
    GetParametersState,
    SendGetTransientParametersState,
    SetParameterValuesState,
    WaitEmptyMessageState,
    WaitGetObjectParametersState,
    WaitGetParametersState,
    WaitGetTransientParametersState,
    WaitInformMRebootState,
    WaitInformState,
    WaitRebootResponseState,
    WaitSetParameterValuesState,
)
from magma.enodebd.state_machines.enb_acs_impl import EnodebAcsStateMachine
from magma.enodebd.tr069 import models
class BaicellsRTSGetObjectParametersState(EnodebAcsState):
    """
    Get information on parameters belonging to objects that can be added or
    removed from the configuration.

    Baicells QAFB will report a parameter value as None if it does not exist
    in the data model, rather than replying with a Fault message like most
    eNB devices.
    """

    def __init__(
            self,
            acs: EnodebAcsStateMachine,
            when_delete: str,
            when_add: str,
            when_set: str,
            when_skip: str,
    ):
        super().__init__()
        self.acs = acs
        self.rm_obj_transition = when_delete
        self.add_obj_transition = when_add
        self.set_params_transition = when_set
        self.skip_transition = when_skip

    def get_msg(self, message: Any) -> AcsMsgAndTransition:
        """ Respond with GetParameterValuesRequest """
        names = _get_object_params_to_get(
            self.acs.device_cfg,
            self.acs.data_model,
        )

        # Generate the request
        request = models.GetParameterValues()
        request.ParameterNames = models.ParameterNames()
        request.ParameterNames.arrayType = 'xsd:string[%d]' \
                                           % len(names)
        request.ParameterNames.string = []
        for name in names:
            path = self.acs.data_model.get_parameter(name).path
            request.ParameterNames.string.append(path)

        return AcsMsgAndTransition(request, None)

    def read_msg(self, message: Any) -> AcsReadMsgResult:
        """
        Process GetParameterValuesResponse

        Object parameters that have a reported value of None indicate that
        the object is not in the eNB's configuration. Most eNB devices will
        reply with a Fault message if we try to get values of parameters that
        don't exist on the data model, so this is an idiosyncrasy of Baicells
        QAFB.
        """
        if not isinstance(message, models.GetParameterValuesResponse):
            return AcsReadMsgResult(False, None)

        path_to_val = {}
        for param_value_struct in message.ParameterList.ParameterValueStruct:
            path_to_val[param_value_struct.Name] = \
                param_value_struct.Value.Data

        logger.debug('Received object parameters: %s', str(path_to_val))

        num_freq_neighbor = self.acs.data_model.get_num_neighbor_freq()
        for i in range(1, num_freq_neighbor + 1):
            obj_name = ParameterName.NEGIH_FREQ_LIST % i
            obj_to_params = self.acs.data_model.get_numbered_param_names()
            param_name_list = obj_to_params[obj_name]
            for name in param_name_list:
                path = self.acs.data_model.get_parameter(name).path
                if path in path_to_val:
                    value = path_to_val[path]
                    if value is None:
                        continue
                    if obj_name not in self.acs.device_cfg.get_object_names():
                        self.acs.device_cfg.add_object(obj_name)
                    magma_value = \
                        self.acs.data_model.transform_for_magma(name, value)
                    self.acs.device_cfg.set_parameter_for_object(
                        name,
                        magma_value,
                        obj_name,
                    )

        # Now we have enough information to build the desired configuration
        if self.acs.desired_cfg is None:
            self.acs.desired_cfg = build_desired_config(
                self.acs.mconfig,
                self.acs.service_config,
                self.acs.device_cfg,
                self.acs.data_model,
                self.acs.config_postprocessor,
            )

        if len(
            get_all_objects_to_delete(
                self.acs.desired_cfg,
                self.acs.device_cfg,
            ),
        ) > 0:
            return AcsReadMsgResult(True, self.rm_obj_transition)
        elif len(
            get_all_objects_to_add(
                self.acs.desired_cfg,
                self.acs.device_cfg,
            ),
        ) > 0:
            return AcsReadMsgResult(True, self.add_obj_transition)
        elif len(
            get_all_param_values_to_set(
                self.acs.desired_cfg,
                self.acs.device_cfg,
                self.acs.data_model,
            ),
        ) > 0:
            return AcsReadMsgResult(True, self.set_params_transition)
        return AcsReadMsgResult(True, self.skip_transition)

    def state_description(self) -> str:
        return 'Getting object parameters'
def _get_object_params_to_get(
    device_cfg: EnodebConfiguration,
    data_model: DataModel,
) -> List[ParameterName]:
    """
    Returns a list of parameter names for object parameters we don't know the
    current value of.

    Since there is no parameter for tracking the number of PLMNs, then we
    make the assumption that if any PLMN object exists, then we've already
    fetched the object parameter values.
    """
    if device_cfg.has_object(ParameterName.NEGIH_FREQ_LIST % 1) :
        return []

    names = []
    num_freq_neighbor = data_model.get_num_neighbor_freq()
    obj_to_params = data_model.get_numbered_param_names()
    for i in range(1, num_freq_neighbor + 1):
        obj_name = ParameterName.NEGIH_FREQ_LIST % i
        desired = obj_to_params[obj_name]
        names = names + desired
    return names

class BaicellsRTSWaitGetTransientParametersState(EnodebAcsState):
    """
    Periodically read eNodeB status. Note: keep frequency low to avoid
    backing up large numbers of read operations if enodebd is busy
    """

    def __init__(
            self,
            acs: EnodebAcsStateMachine,
            when_get: str,
            when_get_obj_params: str,
            when_delete: str,
            when_add: str,
            when_set: str,
            when_skip: str,
    ):
        super().__init__()
        self.acs = acs
        self.done_transition = when_get
        self.get_obj_params_transition = when_get_obj_params
        self.rm_obj_transition = when_delete
        self.add_obj_transition = when_add
        self.set_transition = when_set
        self.skip_transition = when_skip

    def read_msg(self, message: Any) -> AcsReadMsgResult:
        """ Process GetParameterValuesResponse """
        if not isinstance(message, models.GetParameterValuesResponse):
            return AcsReadMsgResult(False, None)
        # Current values of the fetched parameters
        name_to_val = parse_get_parameter_values_response(
            self.acs.data_model,
            message,
        )
        logger.debug('Received Parameters: %s', str(name_to_val))

        # Update device configuration
        for name in name_to_val:
            magma_value = \
                self.acs.data_model.transform_for_magma(name, name_to_val[name])
            self.acs.device_cfg.set_parameter(name, magma_value)

        return AcsReadMsgResult(True, self.get_next_state())

    def get_next_state(self) -> str:
        should_get_params = \
            len(
                get_params_to_get(
                    self.acs.device_cfg,
                    self.acs.data_model,
                ),
            ) > 0
        if should_get_params:
            return self.done_transition
        should_get_obj_params = \
            len(
                _get_object_params_to_get(
                    self.acs.device_cfg,
                    self.acs.data_model,
                ),
            ) > 0
        if should_get_obj_params:
            return self.get_obj_params_transition
        elif len(
            get_all_objects_to_delete(
                self.acs.desired_cfg,
                self.acs.device_cfg,
            ),
        ) > 0:
            return self.rm_obj_transition
        elif len(
            get_all_objects_to_add(
                self.acs.desired_cfg,
                self.acs.device_cfg,
            ),
        ) > 0:
            return self.add_obj_transition
        return self.skip_transition

    def state_description(self) -> str:
        return 'Getting transient read-only parameters'


class BaicellsRTSHandler(BasicEnodebAcsStateMachine):
    def __init__(
            self,
            service: MagmaService,
    ) -> None:
        self._state_map = {}
        super().__init__(service)

    def reboot_asap(self) -> None:
        self.transition('reboot')

    def is_enodeb_connected(self) -> bool:
        return not isinstance(self.state, WaitInformState)

    def _init_state_map(self) -> None:
        self._state_map = {
            'wait_inform': WaitInformState(self, when_done='wait_empty', when_boot=None),
            'wait_empty': WaitEmptyMessageState(self, when_done='get_transient_params', when_missing='check_optional_params'),
            'check_optional_params': CheckOptionalParamsState(self, when_done='get_transient_params'),
            'get_transient_params': SendGetTransientParametersState(self, when_done='wait_get_transient_params'),
            'wait_get_transient_params': WaitGetTransientParametersState(self, when_get='get_params', when_get_obj_params='get_obj_params', when_delete='delete_objs', when_add='add_objs', when_set='set_params', when_skip='end_session'),
            'get_params': GetParametersState(self, when_done='wait_get_params'),
            'wait_get_params': WaitGetParametersState(self, when_done='get_obj_params'),
            'get_obj_params': GetObjectParametersState(self, when_done='wait_get_obj_params'),
            'wait_get_obj_params': WaitGetObjectParametersState(self, when_delete='delete_objs', when_add='add_objs', when_set='set_params', when_skip='end_session'),
            'delete_objs': DeleteObjectsState(self, when_add='add_objs', when_skip='set_params'),
            'add_objs': AddObjectsState(self, when_done='set_params'),
            'set_params': SetParameterValuesState(self, when_done='wait_set_params'),
            'wait_set_params': WaitSetParameterValuesState(self, when_done='check_get_params', when_apply_invasive='reboot'),
            'check_get_params': GetParametersState(self, when_done='check_wait_get_params', request_all_params=True),
            'check_wait_get_params': WaitGetParametersState(self, when_done='end_session'),
            'end_session': EndSessionState(self),
            'reboot': BaicellsSendRebootState(self, when_done='wait_reboot'),
            'wait_reboot': WaitRebootResponseState(self, when_done='wait_post_reboot_inform'),
            'wait_post_reboot_inform': WaitInformMRebootState(self, when_done='wait_empty_post_reboot', when_timeout='wait_inform_post_reboot'),
            'wait_inform_post_reboot': WaitInformState(self, when_done='wait_empty_post_reboot', when_boot=None),
            'wait_empty_post_reboot': WaitEmptyMessageState(self, when_done='get_transient_params', when_missing='check_optional_params'),
            # The states below are entered when an unexpected message type is
            # received
            'unexpected_fault': ErrorState(self),
        }

    @property
    def device_name(self) -> str:
        return EnodebDeviceName.BAICELLS_RTS

    @property
    def data_model_class(self) -> Type[DataModel]:
        return BaicellsRTSTrDataModel

    @property
    def config_postprocessor(self) -> EnodebConfigurationPostProcessor:
        return BaicellsRTSTrConfigurationInitializer()

    @property
    def state_map(self) -> Dict[str, EnodebAcsState]:
        return self._state_map

    @property
    def disconnected_state_name(self) -> str:
        return 'wait_inform'

    @property
    def unexpected_fault_state_name(self) -> str:
        return 'unexpected_fault'


class BaicellsRTSTrDataModel(DataModel):
    """
    Class to represent relevant data model parameters from TR-196/TR-098/TR-181.
    This class is effectively read-only

    This is for any version beginning with BaiBS_ or after
    """
    # Parameters to query when reading eNodeB config
    LOAD_PARAMETERS = [ParameterName.DEVICE]
    # Mapping of TR parameter paths to aliases
    DEVICE_PATH = 'Device.'
    FAPSERVICE_PATH = DEVICE_PATH + 'Services.FAPService.1.'
    PARAMETERS = {
        # Top-level objects
        ParameterName.DEVICE: TrParam(DEVICE_PATH, True, TrParameterType.OBJECT, False),
        ParameterName.FAP_SERVICE: TrParam(FAPSERVICE_PATH, True, TrParameterType.OBJECT, False),

        # Device info parameters
        ParameterName.GPS_STATUS: TrParam(DEVICE_PATH + 'DeviceInfo.X_BAICELLS_COM_GPS_Status', True, TrParameterType.BOOLEAN, False),
        ParameterName.PTP_STATUS: TrParam(DEVICE_PATH + 'DeviceInfo.X_BAICELLS_COM_1588_Status', True, TrParameterType.BOOLEAN, False),
        ParameterName.MME_STATUS: TrParam(DEVICE_PATH + 'DeviceInfo.X_BAICELLS_COM_MME_Status', True, TrParameterType.BOOLEAN, False),
        ParameterName.REM_STATUS: TrParam(FAPSERVICE_PATH + 'REM.X_BAICELLS_COM_REM_Status', True, TrParameterType.BOOLEAN, False),
        ParameterName.LOCAL_GATEWAY_ENABLE:
            TrParam(DEVICE_PATH + 'DeviceInfo.X_BAICELLS_COM_LTE_LGW_Switch', False, TrParameterType.BOOLEAN, False),
        # Tested Baicells devices were missing this parameter
        ParameterName.GPS_ENABLE: TrParam(DEVICE_PATH + 'X_BAICELLS_COM_GpsSyncEnable', False, TrParameterType.BOOLEAN, True),
        ParameterName.GPS_LAT: TrParam(DEVICE_PATH + 'FAP.GPS.LockedLatitude', True, TrParameterType.INT, True),
        ParameterName.GPS_LONG: TrParam(DEVICE_PATH + 'FAP.GPS.LockedLongitude', True, TrParameterType.INT, True),
        ParameterName.SW_VERSION: TrParam(DEVICE_PATH + 'DeviceInfo.SoftwareVersion', True, TrParameterType.STRING, False),
        ParameterName.SERIAL_NUMBER: TrParam(DEVICE_PATH + 'DeviceInfo.SerialNumber', True, TrParameterType.STRING, False),

        # Capabilities
        ParameterName.DUPLEX_MODE_CAPABILITY: TrParam(
            FAPSERVICE_PATH + 'Capabilities.LTE.DuplexMode', True, TrParameterType.STRING, False,
        ),
        ParameterName.BAND_CAPABILITY: TrParam(FAPSERVICE_PATH + 'Capabilities.LTE.BandsSupported', True, TrParameterType.STRING, False),

        # RF-related parameters
        ParameterName.EARFCNDL: TrParam(FAPSERVICE_PATH + 'X_BAICELLS_COM_LTE.EARFCNDLInUse', True, TrParameterType.INT, False),
        ParameterName.EARFCNUL: TrParam(FAPSERVICE_PATH + 'X_BAICELLS_COM_LTE.EARFCNULInUse', True, TrParameterType.INT, False),
        ParameterName.BAND: TrParam(FAPSERVICE_PATH + 'CellConfig.LTE.RAN.RF.FreqBandIndicator', True, TrParameterType.INT, False),
        ParameterName.PCI: TrParam(FAPSERVICE_PATH + 'CellConfig.LTE.RAN.RF.PhyCellID', False, TrParameterType.INT, False),
        ParameterName.DL_BANDWIDTH: TrParam(FAPSERVICE_PATH + 'CellConfig.LTE.RAN.RF.DLBandwidth', True, TrParameterType.STRING, False),
        ParameterName.UL_BANDWIDTH: TrParam(FAPSERVICE_PATH + 'CellConfig.LTE.RAN.RF.ULBandwidth', True, TrParameterType.STRING, False),
        ParameterName.SUBFRAME_ASSIGNMENT: TrParam(
            FAPSERVICE_PATH
            + 'CellConfig.LTE.RAN.PHY.TDDFrame.SubFrameAssignment', True, TrParameterType.INT, False,
        ),
        ParameterName.SPECIAL_SUBFRAME_PATTERN: TrParam(
            FAPSERVICE_PATH
            + 'CellConfig.LTE.RAN.PHY.TDDFrame.SpecialSubframePatterns', True, TrParameterType.INT, False,
        ),
        ParameterName.CELL_ID: TrParam(FAPSERVICE_PATH + 'CellConfig.LTE.RAN.Common.CellIdentity', True, TrParameterType.UNSIGNED_INT, False),

        # Other LTE parameters
        ParameterName.ADMIN_STATE: TrParam(FAPSERVICE_PATH + 'FAPControl.LTE.AdminState', False, TrParameterType.BOOLEAN, False),
        ParameterName.OP_STATE: TrParam(FAPSERVICE_PATH + 'FAPControl.LTE.OpState', True, TrParameterType.BOOLEAN, False),
        ParameterName.RF_TX_STATUS: TrParam(FAPSERVICE_PATH + 'FAPControl.LTE.RFTxStatus', True, TrParameterType.BOOLEAN, False),

        # RAN parameters
        ParameterName.CELL_RESERVED: TrParam(
            FAPSERVICE_PATH
            + 'CellConfig.LTE.RAN.CellRestriction.CellReservedForOperatorUse', True, TrParameterType.BOOLEAN, False,
        ),
        ParameterName.CELL_BARRED: TrParam(
            FAPSERVICE_PATH
            + 'CellConfig.LTE.RAN.CellRestriction.CellBarred', True, TrParameterType.BOOLEAN, False,
        ),

        # Core network parameters
        ParameterName.MME_IP: TrParam(
            FAPSERVICE_PATH + 'FAPControl.LTE.Gateway.S1SigLinkServerList', True, TrParameterType.STRING, False,
        ),
        ParameterName.MME_PORT: TrParam(FAPSERVICE_PATH + 'FAPControl.LTE.Gateway.S1SigLinkPort', True, TrParameterType.INT, False),
        ParameterName.NUM_PLMNS: TrParam(
            FAPSERVICE_PATH + 'CellConfig.LTE.EPC.PLMNListNumberOfEntries', True, TrParameterType.INT, False,
        ),
        ParameterName.PLMN: TrParam(FAPSERVICE_PATH + 'CellConfig.LTE.EPC.PLMNList.', True, TrParameterType.STRING, False),
        # PLMN arrays are added below
        ParameterName.TAC: TrParam(FAPSERVICE_PATH + 'CellConfig.LTE.EPC.TAC', True, TrParameterType.INT, False),
        ParameterName.IP_SEC_ENABLE: TrParam(
            DEVICE_PATH + 'Services.FAPService.Ipsec.IPSEC_ENABLE', False, TrParameterType.BOOLEAN, False,
        ),
        ParameterName.MME_POOL_ENABLE: TrParam(
            FAPSERVICE_PATH
            + 'FAPControl.LTE.Gateway.X_BAICELLS_COM_MmePool.Enable', True, TrParameterType.BOOLEAN, False,
        ),

        # Management server parameters
        ParameterName.PERIODIC_INFORM_ENABLE:
            TrParam(DEVICE_PATH + 'ManagementServer.PeriodicInformEnable', False, TrParameterType.BOOLEAN, False),
        ParameterName.PERIODIC_INFORM_INTERVAL:
            TrParam(DEVICE_PATH + 'ManagementServer.PeriodicInformInterval', False, TrParameterType.INT, False),

        # Performance management parameters
        ParameterName.PERF_MGMT_ENABLE: TrParam(
            DEVICE_PATH + 'FAP.PerfMgmt.Config.1.Enable', False, TrParameterType.BOOLEAN, False,
        ),
        ParameterName.PERF_MGMT_UPLOAD_INTERVAL: TrParam(
            DEVICE_PATH + 'FAP.PerfMgmt.Config.1.PeriodicUploadInterval', False, TrParameterType.INT, False,
        ),
        ParameterName.PERF_MGMT_UPLOAD_URL: TrParam(
            DEVICE_PATH + 'FAP.PerfMgmt.Config.1.URL', False, TrParameterType.STRING, False,
        ),

    }

    NUM_PLMNS_IN_CONFIG = 6
    for i in range(1, NUM_PLMNS_IN_CONFIG + 1):
        PARAMETERS[(ParameterName.PLMN_N) % i] = TrParam(
            FAPSERVICE_PATH + 'CellConfig.LTE.EPC.PLMNList.%d.' % i, True, TrParameterType.STRING, False,
        )
        PARAMETERS[ParameterName.PLMN_N_CELL_RESERVED % i] = TrParam(
            FAPSERVICE_PATH
            + 'CellConfig.LTE.EPC.PLMNList.%d.CellReservedForOperatorUse' % i, True, TrParameterType.BOOLEAN, False,
        )
        PARAMETERS[ParameterName.PLMN_N_ENABLE % i] = TrParam(
            FAPSERVICE_PATH + 'CellConfig.LTE.EPC.PLMNList.%d.Enable' % i, True, TrParameterType.BOOLEAN, False,
        )
        PARAMETERS[ParameterName.PLMN_N_PRIMARY % i] = TrParam(
            FAPSERVICE_PATH + 'CellConfig.LTE.EPC.PLMNList.%d.IsPrimary' % i, True, TrParameterType.BOOLEAN, False,
        )
        PARAMETERS[ParameterName.PLMN_N_PLMNID % i] = TrParam(
            FAPSERVICE_PATH + 'CellConfig.LTE.EPC.PLMNList.%d.PLMNID' % i, True, TrParameterType.STRING, False,
        )
    NUM_NEIGHBOR_CELL_CONFIG = 16
    NUM_NEIGHBOR_FREQ_CONFIG = 8
    for i in range(1, NUM_NEIGHBOR_FREQ_CONFIG + 1):
        PARAMETERS[(ParameterName.NEGIH_FREQ_LIST) % i] = TrParam(
            FAPSERVICE_PATH + 'CellConfig.LTE.RAN.Mobility.IdleMode.InterFreq.Carrier.%d.' % i, True, TrParameterType.UNSIGNED_INT, False,
        )
        PARAMETERS[ParameterName.NEIGHBOR_FREQ_INDEX_N % i] = TrParam(
            FAPSERVICE_PATH + 'CellConfig.LTE.RAN.Mobility.IdleMode.InterFreq.Carrier.%d.' % i, True, TrParameterType.UNSIGNED_INT, False,
        )
        PARAMETERS[ParameterName.NEIGHBOR_FREQ_EARFCN_N % i] = TrParam(
            FAPSERVICE_PATH + 'CellConfig.LTE.RAN.Mobility.IdleMode.InterFreq.Carrier.%d.EUTRACarrierARFCN' % i, True,
            TrParameterType.UNSIGNED_INT, False,
        )
        PARAMETERS[ParameterName.NEIGHBOR_FREQ_QRXLEVMINSIB5_N % i] = TrParam(
            FAPSERVICE_PATH + 'CellConfig.LTE.RAN.Mobility.IdleMode.InterFreq.Carrier.%d.QRxLevMinSIB5' % i, True,
            TrParameterType.INT, False,
        )
        PARAMETERS[ParameterName.NEIGHBOR_FREQ_Q_OFFSETRANGE_N % i] = TrParam(
            FAPSERVICE_PATH + 'CellConfig.LTE.RAN.Mobility.IdleMode.InterFreq.Carrier.%d.QOffsetFreq' % i, True,
            TrParameterType.INT, False,
        )
        PARAMETERS[ParameterName.NEIGHBOR_FREQ_TRESELECTIONEUTRA_N % i] = TrParam(
            FAPSERVICE_PATH + 'CellConfig.LTE.RAN.Mobility.IdleMode.InterFreq.Carrier.%d.TReselectionEUTRA' % i, True,
            TrParameterType.UNSIGNED_INT, False,
        )
        PARAMETERS[ParameterName.NEIGHBOR_FREQ_RESELECTIONPRIORITY_N % i] = TrParam(
            FAPSERVICE_PATH + 'CellConfig.LTE.RAN.Mobility.IdleMode.InterFreq.Carrier.%d.CellReselectionPriority' % i, True,
            TrParameterType.UNSIGNED_INT, False,
        )
        PARAMETERS[ParameterName.NEIGHBOR_FREQ_RESELTHRESHHIGH_N % i] = TrParam(
            FAPSERVICE_PATH + 'CellConfig.LTE.RAN.Mobility.IdleMode.InterFreq.Carrier.%d.ThreshXHigh' % i, True,
            TrParameterType.UNSIGNED_INT, False,
        )
        PARAMETERS[ParameterName.NEIGHBOR_FREQ_RESELTHRESHLOW_N % i] = TrParam(
            FAPSERVICE_PATH + 'CellConfig.LTE.RAN.Mobility.IdleMode.InterFreq.Carrier.%d.ThreshXLow' % i, True,
            TrParameterType.UNSIGNED_INT, False,
        )
        PARAMETERS[ParameterName.NEIGHBOR_FREQ_PMAX_N % i] = TrParam(
            FAPSERVICE_PATH + 'CellConfig.LTE.RAN.Mobility.IdleMode.InterFreq.Carrier.%d.PMax' % i, True,
            TrParameterType.INT, False,
        )
        PARAMETERS[ParameterName.NEIGHBOR_FREQ_TRESELECTIONEUTRASFMEDIUM_N % i] = TrParam(
            FAPSERVICE_PATH + 'CellConfig.LTE.RAN.Mobility.IdleMode.InterFreq.Carrier.%d.TReselectionEUTRASFMedium' % i, True,
            TrParameterType.UNSIGNED_INT, False,
        )
        PARAMETERS[ParameterName.NEIGHBOR_FREQ_ENABLE_N % i] = TrParam(
            FAPSERVICE_PATH + 'CellConfig.LTE.RAN.Mobility.IdleMode.InterFreq.Carrier.%d.Enable' % i, True, TrParameterType.BOOLEAN, False,

        )

    TRANSFORMS_FOR_ENB = {
        ParameterName.DL_BANDWIDTH: transform_for_enb.bandwidth,
        ParameterName.UL_BANDWIDTH: transform_for_enb.bandwidth,
    }
    TRANSFORMS_FOR_MAGMA = {
        ParameterName.DL_BANDWIDTH: transform_for_magma.bandwidth,
        ParameterName.UL_BANDWIDTH: transform_for_magma.bandwidth,
        # We don't set GPS, so we don't need transform for enb
        ParameterName.GPS_LAT: transform_for_magma.gps_tr181,
        ParameterName.GPS_LONG: transform_for_magma.gps_tr181,
    }

    @classmethod
    def get_parameter(cls, param_name: ParameterName) -> Optional[TrParam]:
        return cls.PARAMETERS.get(param_name)

    @classmethod
    def _get_magma_transforms(
            cls,
    ) -> Dict[ParameterName, Callable[[Any], Any]]:
        return cls.TRANSFORMS_FOR_MAGMA

    @classmethod
    def _get_enb_transforms(cls) -> Dict[ParameterName, Callable[[Any], Any]]:
        return cls.TRANSFORMS_FOR_ENB

    @classmethod
    def get_load_parameters(cls) -> List[ParameterName]:
        return cls.LOAD_PARAMETERS

    @classmethod
    def get_num_plmns(cls) -> int:
        return cls.NUM_PLMNS_IN_CONFIG
    @classmethod
    def get_num_neighbor_freq(cls) -> int:
        return cls.NUM_NEIGHBOR_FREQ_CONFIG
    @classmethod
    def get_num_neighbor_cell(cls) -> int:
        return cls.NUM_NEIGHBOR_CELL_CONFIG
    @classmethod
    def get_parameter_names(cls) -> List[ParameterName]:
        excluded_params = [
            str(ParameterName.DEVICE),
            str(ParameterName.FAP_SERVICE),
        ]
        names = list(
            filter(
                lambda x: (not str(x).startswith('PLMN'))
                          and (str(x) not in excluded_params),
                cls.PARAMETERS.keys(),
            ),
        )
        return names

    @classmethod
    def get_numbered_param_names(cls) -> Dict[ParameterName, List[ParameterName]]:
        names = {}
        for i in range(1, cls.NUM_PLMNS_IN_CONFIG + 1):
            params = []
            params.append(ParameterName.PLMN_N_CELL_RESERVED % i)
            params.append(ParameterName.PLMN_N_ENABLE % i)
            params.append(ParameterName.PLMN_N_PRIMARY % i)
            params.append(ParameterName.PLMN_N_PLMNID % i)
            names[ParameterName.PLMN_N % i] = params
        for i in range(1,cls.NUM_NEIGHBOR_FREQ_CONFIG + 1):
            params = []
            params.append(ParameterName.NEIGHBOR_FREQ_ENABLE_N % i)
            params.append(ParameterName.NEIGHBOR_FREQ_INDEX_N % i)
            params.append(ParameterName.NEIGHBOR_FREQ_EARFCN_N % i)
            params.append(ParameterName.NEIGHBOR_FREQ_PMAX_N % i)
            params.append(ParameterName.NEIGHBOR_FREQ_Q_OFFSETRANGE_N % i)
            params.append(ParameterName.NEIGHBOR_FREQ_RESELTHRESHLOW_N % i)
            params.append(ParameterName.NEIGHBOR_FREQ_RESELTHRESHHIGH_N % i)
            params.append(ParameterName.NEIGHBOR_FREQ_RESELECTIONPRIORITY_N % i)
            params.append(ParameterName.NEIGHBOR_FREQ_TRESELECTIONEUTRASFMEDIUM_N % i)
            params.append(ParameterName.NEIGHBOR_FREQ_QRXLEVMINSIB5_N)
            params.append(ParameterName.NEIGHBOR_FREQ_TRESELECTIONEUTRA_N % i)
            names[ParameterName.NEGIH_FREQ_LIST % i] = params
        return names


class BaicellsRTSTrConfigurationInitializer(EnodebConfigurationPostProcessor):
    def postprocess(self, desired_cfg: EnodebConfiguration) -> None:
        desired_cfg.set_parameter(ParameterName.CELL_BARRED, False)

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


class ParameterName():
    # Top-level objects
    DEVICE = 'Device'
    FAP_SERVICE = 'FAPService'

    # Device info parameters
    GPS_STATUS = 'GPS status'
    PTP_STATUS = 'PTP status'
    MME_STATUS = 'MME status'
    REM_STATUS = 'REM status'
    RF_STATE = 'RF state'
    SYNC_1588_STATUS = '1588 Sync status'
    LATITUDE = 'Latitude'
    LONGITUDE = 'Longitude'
    ALTITUDE = 'Altitude'
    UPTIME = 'Uptime'

    LOCAL_GATEWAY_ENABLE = 'Local gateway enable'
    GPS_ENABLE = 'GPS enable'
    GPS_LAT = 'GPS lat'
    GPS_LONG = 'GPS long'
    SW_VERSION = 'SW version'

    SERIAL_NUMBER = 'Serial number'
    VENDOR = 'Vendor'
    MODEL_NAME = 'Model Name'
    IP_ADDRESS = 'Ip Address'

    CELL_ID = 'Cell ID'

    # Capabilities
    DUPLEX_MODE_CAPABILITY = 'Duplex mode capability'
    BAND_CAPABILITY = 'Band capability'

    # RF-related parameters
    EARFCNDL = 'EARFCNDL'
    EARFCNUL = 'EARFCNUL'
    BAND = 'Band'
    PCI = 'PCI'
    DL_BANDWIDTH = 'DL bandwidth'
    UL_BANDWIDTH = 'UL bandwidth'
    SUBFRAME_ASSIGNMENT = 'Subframe assignment'
    SPECIAL_SUBFRAME_PATTERN = 'Special subframe pattern'

    # Other LTE parameters
    ADMIN_STATE = 'Admin state'
    OP_STATE = 'Opstate'
    RF_TX_STATUS = 'RF TX status'

    # RAN parameters
    CELL_RESERVED = 'Cell reserved'
    CELL_BARRED = 'Cell barred'

    # Core network parameters
    MME_IP = 'MME IP'
    MME_PORT = 'MME port'
    NUM_PLMNS = 'Num PLMNs'
    PLMN = 'PLMN'
    PLMN_LIST = 'PLMN List'
    MME_POOL_1 = 'MME Pool 1'
    MME_POOL_2 = 'MME Pool 2'

    # PLMN parameters
    PLMN_N = 'PLMN %d'
    PLMN_N_CELL_RESERVED = 'PLMN %d cell reserved'
    PLMN_N_ENABLE = 'PLMN %d enable'
    PLMN_N_PRIMARY = 'PLMN %d primary'
    PLMN_N_PLMNID = 'PLMN %d PLMNID'

    # PLMN arrays are added below
    TAC = 'TAC'
    IP_SEC_ENABLE = 'IPSec enable'
    MME_POOL_ENABLE = 'MME pool enable'

    # Management server parameters
    PERIODIC_INFORM_ENABLE = 'Periodic inform enable'
    PERIODIC_INFORM_INTERVAL = 'Periodic inform interval'

    # Performance management parameters
    PERF_MGMT_ENABLE = 'Perf mgmt enable'
    PERF_MGMT_UPLOAD_INTERVAL = 'Perf mgmt upload interval'
    PERF_MGMT_UPLOAD_URL = 'Perf mgmt upload URL'
    PERF_MGMT_USER = 'Perf mgmt username'
    PERF_MGMT_PASSWORD = 'Perf mgmt password'

    #Power control parameters
    REFERENCE_SIGNAL_POWER = 'Reference Signal Power'
    POWER_CLASS = 'PowerClass'
    PA = 'PA'
    PB = 'PB'

    #IP Configuration
    #IP_ACCESS_MODE = 'IP Access Mode'
    #NET_MASK = 'net mask'
    MTU = 'mtu'
    #HOST_NAME = 'Host Name'
    TIME_ZONE = 'TimeZone'
    #DNS_ADDRESS_1 = 'DNS Address 1'
    #DNS_ADDRESS_2 = 'DNS Address 2'
    #DNS_ADDRESS_3 = 'DNS Address 3'

    #HO algorithm
    A1_THRESHOLD_RSRP = 'A1 Threshold - RSRPc'
    #LTE_A1_THRESHOLD_RSRQ = 'A1 Threshold - RSRP'
    #A1_HYSTERESIS = 'A1 Hysteresisc'
    #A1_TIME_TO_TRIGGER = 'Time To Trigger'

    A2_THRESHOLD_RSRP = 'A2 Threshold - RSRP'
    #LTE_A2_THRESHOLD_RSRQ = 'LTE_A2_THRESHOLD_RSRQ'
    #LTE_A2_THRESHOLD_RSRP_IRAT_VOLTE = 'LTE_A2_THRESHOLD_RSRP_IRAT_VOLTE'
    #LTE_A2_THRESHOLD_RSRQ_IRAT_VOLTE = 'LTE_A2_THRESHOLD_RSRQ_IRAT_VOLTE'
    #LTE_A2_THRESHOLD_RSRP_IRAT_DATA = 'LTE_A2_THRESHOLD_RSRP_IRAT_DATA'
    #LTE_A2_THRESHOLD_RSRQ_IRAT_DATA = 'LTE_A2_THRESHOLD_RSRQ_IRAT_DATA'
    #LTE_A2_THRESHOLD_RSRP_BLIND_DIR = 'LTE_A2_THRESHOLD_RSRP_BLIND_DIR'
    #LTE_A2_THRESHOLD_RSRQ_BLIND_DIR = 'LTE_A2_THRESHOLD_RSRQ_BLIND_DIR'
    #A2_HYSTERESIS = 'A2 Hysteresis'
    #A2_TIME_TO_TRIGGER = 'A2 Time_To_Trigger'

    A3_OFFSET = 'A3 Offset'
    A3_OFFSET_ANR = 'A3_Offset_ANR'
    #LTE_ANR_RSRP_THRESHOLD = 'LTE_ANR_RSRP_THRESHOLD'
    #A3_HYSTERESIS = 'A3 Hysteresis'
    #A3_TIME_TO_TRIGGER = 'A3 Time_To_Trigger'

    A4_THRESHOLD_RSRP = 'A4 Threshold_RSRP'
    #LTE_A4_THRESHOLD_RSRQ = 'LTE_A4_THRESHOLD_RSRQ'

    LTE_INTRA_A5_THRESHOLD_1_RSRP = 'LTE_INTRA_A5_THRESHOLD_1_RSRP'
    LTE_INTRA_A5_THRESHOLD_2_RSRP = 'LTE_INTRA_A5_THRESHOLD_2_RSRP'
    #LTE_INTER_A5_THRESHOLD_1_RSRP = 'LTE_INTER_A5_THRESHOLD_1_RSRP'
    #LTE_INTER_A5_THRESHOLD_2_RSRP = 'LTE_INTER_A5_THRESHOLD_2_RSRP'
    #LTE_A5_THRESHOLD_1_RSRQ = 'LTE_A5_THRESHOLD_1_RSRQ'
    #LTE_A5_THRESHOLD_2_RSRQ = 'LTE_A5_THRESHOLD_2_RSRQ'
    #LTE_INTER_ANR_A5_THRESHOLD_1_RSRP = 'LTE_INTER_ANR_A5_THRESHOLD_1_RSRP'
    #LTE_INTER_ANR_A5_THRESHOLD_2_RSRP = 'LTE_INTER_ANR_A5_THRESHOLD_2_RSRP'
    #A5_HYSTERESIS = 'A5_Hysteresis'
    #A5_TIME_TO_TRIGGER = 'A5 Time_To_Trigger'

    B2_THRESHOLD1_RSRP = 'B2_Threshold1_RSRP'
    B2_THRESHOLD2_RSRP = 'B2_Threshold2_RSRP'
    B2_GERAN_IRAT_THRESHOLD = 'B2_GERAN_IRAT_Threshold'
    #B2_HYSTERESIS  = 'B2 Hysteresis'
    #B2_TIME_TO_TRIGGER = 'B2 Time_To_Trigger'

    CELL_SELECTION_PARAMETERS_QRXLEVMIN = 'Cell_selection_parameters_Qrxlevmin'
    CELL_SELECTION_PARAMETERS_QRXLEVMINOFFSET = 'Cell_selection_parameters_QrxlevminoffsetQrxlevminoffset'

    CELL_RESELECTION_PARAMETERS_S_INTRASEARCH = 'Cell_reselection_parameters_S_IntraSearch'
    CELL_RESELECTION_PARAMETERS_S_NONINTRASEARCH = 'Cell_reselection_parameters_S_NonIntraSearch'
    CELL_RESELECTION_PARAMETERS_QRXLEVMIN = 'Cell_reselection_parameters_Qrxlevmin'
    CELL_RESELECTION_PARAMETERS_RESELECTION_PRIORITY = 'Cell_reselection_parameters_Reselection_Priority'
    CELL_RESELECTION_PARAMETERS_THRESHSERVINGLOW = 'Cell_reselection_parameters_ThreshServingLow'
    X2_ENABLE = 'X2 enable/disable'
    CIPHERING_ALGORITHM = 'Ciphering_Algorithm'
    INTEGRITY_ALGORITHM = 'Integrity_Algorithm'

    #management server
    MANAGEMENT_SERVER = 'Management Server'
    MANAGEMENT_SERVER_PORT = 'Management Server Port'
    MANAGEMENT_SERVER_SSL_ENABLE = 'Management Server SSL Enable'

    #Sync
    SYNC_1588_DOMAIN = '1588 domain'
    SYNC_1588_SYNC_MESSAGE = '1588 sync message'
    SYNC_1588_DELAY_REQUEST = '1588 delay request'
    SYNC_1588_HOLDOVER = '1588 holdover'
    SYNC_1588_ASYMMETRY = '1588 asymmetry'
    SYNC_1588_UNICAST_ENABLE = '1588 unicast enable'
    SYNC_1588_UNICAST_SERVERIP = '1588 unicast server IP'


class TrParameterType():
    BOOLEAN = 'boolean'
    STRING = 'string'
    INT = 'int'
    UNSIGNED_INT = 'unsignedInt'
    OBJECT = 'object'

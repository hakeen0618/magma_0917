/**
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @flow strict-local
 * @format
 */
import type {
  enodeb,
  enodeb_configuration,
  network_ran_configs,
  unmanaged_enodeb_configuration,
} from '../../../generated/MagmaAPIBindings';

import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogTitle from '../../theme/design-system/DialogTitle';
import EnodebContext from '../../components/context/EnodebContext';
import FormControl from '@material-ui/core/FormControl';
import FormLabel from '@material-ui/core/FormLabel';
import List from '@material-ui/core/List';
import MenuItem from '@material-ui/core/MenuItem';
import OutlinedInput from '@material-ui/core/OutlinedInput';
import React from 'react';
import Select from '@material-ui/core/Select';
import Switch from '@material-ui/core/Switch';
import Tab from '@material-ui/core/Tab';
import Tabs from '@material-ui/core/Tabs';
import {
  EnodebBandwidthOption,
  EnodebDeviceClass,
} from '../../components/lte/EnodebUtils';

import EnodeConfigEditFdd from './EnodebDetailConfigFdd';
import EnodeConfigEditTdd from './EnodebDetailConfigTdd';

import {AltFormField} from '../../components/FormField';
import {colors, typography} from '../../theme/default';
import {makeStyles} from '@material-ui/styles';
import {useContext, useEffect, useState} from 'react';
import {useEnqueueSnackbar} from '@fbcnms/ui/hooks/useSnackbar';
import {useRouter} from '@fbcnms/ui/hooks';

const CONFIG_TITLE = 'Config';
const RAN_TITLE = 'Ran';
const DEFAULT_ENB_CONFIG = {
  name: '',
  serial: '',
  description: '',
  config: {
    cell_id: 0,
    device_class: 'Baicells Nova-233 G2 OD FDD',
    transmit_enabled: false,
  },
};

const useStyles = makeStyles(_ => ({
  appBarBtn: {
    color: colors.primary.white,
    background: colors.primary.comet,
    fontFamily: typography.button.fontFamily,
    fontWeight: typography.button.fontWeight,
    fontSize: typography.button.fontSize,
    lineHeight: typography.button.lineHeight,
    letterSpacing: typography.button.letterSpacing,

    '&:hover': {
      background: colors.primary.mirage,
    },
  },
  tabBar: {
    backgroundColor: colors.primary.brightGray,
    color: colors.primary.white,
  },
}));

const EditTableType = {
  config: 0,
  ran: 1,
};

type EditProps = {
  editTable: $Keys<typeof EditTableType>,
};

type DialogProps = {
  open: boolean,
  onClose: () => void,
  editProps?: EditProps,
};

type ButtonProps = {
  title: string,
  isLink: boolean,
  editProps?: EditProps,
};

export default function AddEditEnodeButton(props: ButtonProps) {
  const classes = useStyles();
  const [open, setOpen] = useState(false);

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  return (
    <>
      <EnodeEditDialog
        open={open}
        onClose={handleClose}
        editProps={props.editProps}
      />
      {props.isLink ? (
        <Button
          data-testid={(props.editProps?.editTable ?? '') + 'EditButton'}
          component="button"
          variant="text"
          onClick={handleClickOpen}>
          {props.title}
        </Button>
      ) : (
        <Button
          variant="text"
          className={classes.appBarBtn}
          onClick={handleClickOpen}>
          {props.title}
        </Button>
      )}
    </>
  );
}

function EnodeEditDialog(props: DialogProps) {
  const {open, editProps} = props;
  const classes = useStyles();
  const [enb, setEnb] = useState<enodeb>({});
  const {match} = useRouter();
  const ctx = useContext(EnodebContext);
  const enodebSerial: string = match.params.enodebSerial;
  const enbInfo = ctx.state.enbInfo[enodebSerial];
  const lteRanConfigs = ctx.lteRanConfigs;

  const [tabPos, setTabPos] = useState(
    editProps ? EditTableType[editProps.editTable] : 0,
  );

  const onClose = () => {
    // clear existing state
    props.onClose();
  };

  useEffect(() => {
    setTabPos(editProps ? EditTableType[editProps.editTable] : 0);
    setEnb({});
  }, [editProps, open]);

  return (
    <Dialog data-testid="editDialog" open={open} fullWidth={true} maxWidth="sm">
      <DialogTitle
        label={editProps ? 'Edit eNodeB' : 'Add New eNodeB'}
        onClose={onClose}
      />
      <Tabs
        value={tabPos}
        onChange={(_, v) => setTabPos(v)}
        indicatorColor="primary"
        className={classes.tabBar}>
        <Tab key="config" data-testid="configTab" label={CONFIG_TITLE} />; ;
        <Tab
          key="ran"
          data-testid="ranTab"
          disabled={editProps ? false : true}
          label={RAN_TITLE}
        />
      </Tabs>
      {tabPos === 0 && (
        <ConfigEdit
          isAdd={!editProps}
          enb={Object.keys(enb).length != 0 ? enb : enbInfo?.enb}
          lteRanConfigs={lteRanConfigs}
          onClose={onClose}
          onSave={(enb: enodeb) => {
            setEnb(enb);
            if (editProps) {
              onClose();
            } else {
              setTabPos(tabPos + 1);
            }
          }}
        />
      )}
      {tabPos === 1 && (
        <RanEdit
          isAdd={!editProps}
          enb={Object.keys(enb).length != 0 ? enb : enbInfo?.enb}
          lteRanConfigs={lteRanConfigs}
          onClose={onClose}
          onSave={onClose}
        />
      )}
    </Dialog>
  );
}

type Props = {
  isAdd: boolean,
  enb?: enodeb,
  lteRanConfigs: ?network_ran_configs,
  onClose: () => void,
  onSave: enodeb => void,
};

type BandwidthMhzType = $PropertyType<enodeb_configuration, 'bandwidth_mhz'>;

type OptConfig = {
  earfcndl: string,
  bandwidthMhz: BandwidthMhzType,
  specialSubframePattern: string,
  subframeAssignment: string,
  pci: string,
  tac: string,
  reference_signal_power: string,
  power_class: string,
  pa: string,
  pb: string,
  mme_pool_1: string,
  mme_pool_2: string,
  management_server_host: string,
  management_server_port: string,
  management_server_ssl_enable: string,
  sync_1588_asymmetry: string,
  sync_1588_delay_rq_msg_interval: string,
  sync_1588_domain: string,
  sync_1588_holdover: string,
  sync_1588_msg_interval: string,
  sync_1588_unicast_enable: string,
  sync_1588_unicast_serverIp: string,
  sync_1588_switch: string,
};
type OptKey = $Keys<OptConfig>;

export function RanEdit(props: Props) {
  const {match} = useRouter();
  const ctx = useContext(EnodebContext);
  const enodebSerial: string = match.params.enodebSerial;
  const enbInfo = ctx.state.enbInfo[enodebSerial];

  const handleEnbChange = (key: string, val) =>
    setConfig({...config, [key]: val});

  const handleUnmanagedEnbChange = (key: string, val) =>
    setUnmanagedConfig({...unmanagedConfig, [key]: val});

  const handleOptChange = (key: OptKey, val) =>
    setOptConfig({...optConfig, [(key: string)]: val});

  const [error, setError] = useState('');

  const [
    unmanagedConfig,
    setUnmanagedConfig,
  ] = useState<unmanaged_enodeb_configuration>(
    props.enb?.enodeb_config?.unmanaged_config || {
      cell_id: 0,
      ip_address: '',
      tac: 0,
    },
  );

  const [config, setConfig] = useState<enodeb_configuration>(
    props.enb?.enodeb_config?.managed_config || DEFAULT_ENB_CONFIG.config,
  );

  const [enbConfigType, setEnbConfigType] = useState<'MANAGED' | 'UNMANAGED'>(
    props.enb?.enodeb_config?.config_type ?? 'MANAGED',
  );

  const [optConfig, setOptConfig] = useState<OptConfig>({
    earfcndl: String(config.earfcndl ?? ''),
    bandwidthMhz: config.bandwidth_mhz ?? EnodebBandwidthOption['20'],
    specialSubframePattern: String(config.special_subframe_pattern ?? ''),
    subframeAssignment: String(config.subframe_assignment ?? ''),
    pci: String(config.pci ?? ''),
    tac: String(config.tac ?? ''),
    reference_signal_power: String(
      config.radioConfiguration?.powerControlParameters
        ?.reference_signal_power ?? '',
    ),
    power_class: String(
      config.radioConfiguration?.powerControlParameters?.power_class ?? '',
    ),
    pa: String(config.radioConfiguration?.powerControlParameters?.pa ?? ''),
    pb: String(config.radioConfiguration?.powerControlParameters?.pb ?? ''),
    mme_pool_1: String(config.mme_pool_1 ?? ''),
    mme_pool_2: String(config.mme_pool_2 ?? ''),
    management_server_host: String(
      config.managementServer?.management_server_host ?? '',
    ),
    management_server_port: String(
      config.managementServer?.management_server_port ?? '',
    ),
    management_server_ssl_enable:
      config.managementServer?.management_server_ssl_enable ?? false,
    sync_1588_switch: config.sync_1588?.sync_1588_switch ?? false,
    sync_1588_asymmetry: String(config.sync_1588?.sync_1588_asymmetry ?? ''),
    sync_1588_delay_rq_msg_interval: String(
      config.sync_1588?.sync_1588_delay_rq_msg_interval ?? '',
    ),
    sync_1588_domain: String(config.sync_1588?.sync_1588_domain ?? ''),
    sync_1588_holdover: String(config.sync_1588?.sync_1588_holdover ?? ''),
    sync_1588_msg_interval: String(
      config.sync_1588?.sync_1588_msg_interval ?? '',
    ),
    sync_1588_unicast_enable:
      config.sync_1588?.sync_1588_unicast_enable ?? true,
    sync_1588_unicast_serverIp: String(
      config.sync_1588?.sync_1588_unicast_serverIp ?? '',
    ),
  });

  const enqueueSnackbar = useEnqueueSnackbar();

  const onSave = async () => {
    try {
      const enb: enodeb = {
        ...(props.enb || DEFAULT_ENB_CONFIG),
        config:
          enbConfigType === 'MANAGED'
            ? buildRanConfig(config, optConfig)
            : DEFAULT_ENB_CONFIG.config,
        enodeb_config: {
          config_type: enbConfigType,
          managed_config:
            enbConfigType === 'MANAGED'
              ? buildRanConfig(config, optConfig)
              : undefined,
          unmanaged_config:
            enbConfigType === 'UNMANAGED' ? unmanagedConfig : undefined,
        },
      };

      await ctx.setState(enb.serial, {
        enb_state: enbInfo?.enb_state ?? {},
        enb: enb,
      });

      enqueueSnackbar('eNodeb saved successfully', {
        variant: 'success',
      });
      props.onSave(enb);
    } catch (e) {
      setError(e.response?.data?.message ?? e.message);
    }
  };

  return (
    <>
      <DialogContent data-testid="ranEdit">
        <List>
          {error !== '' && (
            <AltFormField label={''}>
              <FormLabel error>{error}</FormLabel>
            </AltFormField>
          )}
          <AltFormField label={'eNodeB Managed Externally'}>
            <Switch
              data-testid="enodeb_config"
              onChange={({target}) =>
                setEnbConfigType(target.checked ? 'UNMANAGED' : 'MANAGED')
              }
              checked={enbConfigType === 'UNMANAGED'}
            />
          </AltFormField>
          {enbConfigType === 'UNMANAGED' ? (
            <>
              <AltFormField label={'Cell ID'}>
                <OutlinedInput
                  data-testid="cellId"
                  type="number"
                  min={0}
                  max={Math.pow(2, 28) - 1}
                  fullWidth={true}
                  value={unmanagedConfig.cell_id}
                  onChange={({target}) =>
                    handleUnmanagedEnbChange('cell_id', parseInt(target.value))
                  }
                />
              </AltFormField>
              <AltFormField label={'TAC'}>
                <OutlinedInput
                  data-testid="tac"
                  type="number"
                  min={0}
                  max={65535}
                  fullWidth={true}
                  value={unmanagedConfig.tac}
                  onChange={({target}) =>
                    handleUnmanagedEnbChange('tac', parseInt(target.value))
                  }
                />
              </AltFormField>
              <AltFormField label={'IP Address'}>
                <OutlinedInput
                  data-testid="ipAddress"
                  fullWidth={true}
                  placeholder="192.168.0.1/24"
                  value={unmanagedConfig.ip_address}
                  onChange={({target}) =>
                    handleUnmanagedEnbChange('ip_address', target.value)
                  }
                />
              </AltFormField>
            </>
          ) : (
            <>
              <AltFormField label={'Device Class'}>
                <FormControl>
                  <Select
                    value={config.device_class}
                    onChange={({target}) =>
                      handleEnbChange(
                        'device_class',
                        coerceValue(target.value, EnodebDeviceClass),
                      )
                    }
                    input={<OutlinedInput id="deviceClass" />}>
                    {Object.keys(EnodebDeviceClass).map(
                      (k: string, idx: number) => (
                        <MenuItem key={idx} value={EnodebDeviceClass[k]}>
                          {EnodebDeviceClass[k]}
                        </MenuItem>
                      ),
                    )}
                  </Select>
                </FormControl>
              </AltFormField>
              <AltFormField label={'Cell ID'}>
                <OutlinedInput
                  data-testid="cellId"
                  type="number"
                  min={0}
                  max={Math.pow(2, 28) - 1}
                  fullWidth={true}
                  value={config.cell_id}
                  onChange={({target}) =>
                    handleEnbChange('cell_id', parseInt(target.value))
                  }
                />
              </AltFormField>
              <AltFormField label={'Bandwidth'}>
                <FormControl>
                  <Select
                    value={optConfig.bandwidthMhz}
                    onChange={({target}) =>
                      handleOptChange(
                        'bandwidthMhz',
                        coerceValue(target.value, EnodebBandwidthOption),
                      )
                    }
                    input={<OutlinedInput id="bandwidth" />}>
                    {Object.keys(EnodebBandwidthOption).map(
                      (k: string, idx: number) => (
                        <MenuItem key={idx} value={EnodebBandwidthOption[k]}>
                          {EnodebBandwidthOption[k]}
                        </MenuItem>
                      ),
                    )}
                  </Select>
                </FormControl>
              </AltFormField>
              {props.lteRanConfigs?.tdd_config && (
                <EnodeConfigEditTdd
                  earfcndl={optConfig.earfcndl}
                  specialSubframePattern={optConfig.specialSubframePattern}
                  subframeAssignment={optConfig.subframeAssignment}
                  setEarfcndl={v => handleOptChange('earfcndl', v)}
                  setSubframeAssignment={v =>
                    handleOptChange('subframeAssignment', v)
                  }
                  setSpecialSubframePattern={v =>
                    handleOptChange('specialSubframePattern', v)
                  }
                />
              )}
              {props.lteRanConfigs?.fdd_config && (
                <EnodeConfigEditFdd
                  earfcndl={optConfig.earfcndl}
                  earfcnul={props.lteRanConfigs.fdd_config.earfcnul.toString()}
                  setEarfcndl={v => handleOptChange('earfcndl', v)}
                />
              )}
              <AltFormField label={'PCI'}>
                <OutlinedInput
                  data-testid="pci"
                  placeholder="Enter PCI"
                  fullWidth={true}
                  value={optConfig.pci}
                  onChange={({target}) => handleOptChange('pci', target.value)}
                />
              </AltFormField>

              <AltFormField label={'TAC'}>
                <OutlinedInput
                  data-testid="tac"
                  placeholder="Enter TAC"
                  fullWidth={true}
                  value={optConfig.tac}
                  onChange={({target}) => handleOptChange('tac', target.value)}
                />
              </AltFormField>
              <AltFormField label={'Reference Signal Power'}>
                <OutlinedInput
                  data-testid="reference_signal_power"
                  placeholder="Enter Reference Signal Power"
                  fullWidth={true}
                  value={optConfig.reference_signal_power}
                  onChange={({target}) =>
                    handleOptChange('reference_signal_power', target.value)
                  }
                />
              </AltFormField>
              <AltFormField label={'Power Class'}>
                <OutlinedInput
                  data-testid="power_class"
                  placeholder="Enter Power Class"
                  fullWidth={true}
                  value={optConfig.power_class}
                  onChange={({target}) =>
                    handleOptChange('power_class', target.value)
                  }
                />
              </AltFormField>
              <AltFormField label={'PA'}>
                <OutlinedInput
                  data-testid="pa"
                  placeholder="Enter PA"
                  fullWidth={true}
                  value={optConfig.pa}
                  onChange={({target}) => handleOptChange('pa', target.value)}
                />
              </AltFormField>
              <AltFormField label={'PB'}>
                <OutlinedInput
                  data-testid="pb"
                  placeholder="Enter PB"
                  fullWidth={true}
                  value={optConfig.pb}
                  onChange={({target}) => handleOptChange('pb', target.value)}
                />
              </AltFormField>
              <AltFormField label={'MME POOL 1'}>
                <OutlinedInput
                  data-testid="mme_pool_1"
                  placeholder="Enter MME POOL 1"
                  fullWidth={true}
                  value={optConfig.mme_pool_1}
                  onChange={({target}) =>
                    handleOptChange('mme_pool_1', target.value)
                  }
                />
              </AltFormField>
              <AltFormField label={'MME POOL 2'}>
                <OutlinedInput
                  data-testid="mme_pool_2"
                  placeholder="Enter MME POOL 2"
                  fullWidth={true}
                  value={optConfig.mme_pool_2}
                  onChange={({target}) =>
                    handleOptChange('mme_pool_2', target.value)
                  }
                />
              </AltFormField>
              <AltFormField label={'Management SSL Enable'}>
                <FormControl variant={'outlined'}>
                  <Select
                    value={optConfig.management_server_ssl_enable ? 1 : 0}
                    onChange={({target}) =>
                      handleOptChange(
                        'management_server_ssl_enable',
                        target.value === 1,
                      )
                    }
                    input={<OutlinedInput id="management_ssl_enable" />}>
                    <MenuItem value={0}>Disabled</MenuItem>
                    <MenuItem value={1}>Enabled</MenuItem>
                  </Select>
                </FormControl>
              </AltFormField>
              <AltFormField label={'Management Server Host'}>
                <OutlinedInput
                  data-testid="management_host"
                  placeholder="Enter server host"
                  fullWidth={true}
                  value={optConfig.management_server_host}
                  onChange={({target}) =>
                    handleOptChange('management_server_host', target.value)
                  }
                />
              </AltFormField>
              <AltFormField label={'Management Server Port'}>
                <OutlinedInput
                  data-testid="management_port"
                  placeholder="Enter Management Server Port"
                  fullWidth={true}
                  value={optConfig.management_server_port}
                  onChange={({target}) =>
                    handleOptChange('management_server_port', target.value)
                  }
                />
              </AltFormField>
              <AltFormField label={'1588 SYNC SWITCH'}>
                <FormControl variant={'outlined'}>
                  <Select
                    value={optConfig.sync_1588_switch ? 1 : 0}
                    onChange={({target}) =>
                      handleOptChange('sync_1588_switch', target.value === 1)
                    }
                    input={<OutlinedInput id="1588_sync_switch" />}>
                    <MenuItem value={0}>Disabled</MenuItem>
                    <MenuItem value={1}>Enabled</MenuItem>
                  </Select>
                </FormControl>
              </AltFormField>
              <AltFormField label={'1588 SYNC Domain Num'}>
                <OutlinedInput
                  data-testid="1588_sync_domain"
                  placeholder="Enter 1588 SYNC Domain Num"
                  fullWidth={true}
                  value={optConfig.sync_1588_domain}
                  onChange={({target}) =>
                    handleOptChange('sync_1588_domain', target.value)
                  }
                />
              </AltFormField>
              <AltFormField label={'1588 SYNC Holdover'}>
                <OutlinedInput
                  data-testid="1588_sync_holdover"
                  placeholder="Enter 1588 SYNC Holdover"
                  fullWidth={true}
                  value={optConfig.sync_1588_holdover}
                  onChange={({target}) =>
                    handleOptChange('sync_1588_holdover', target.value)
                  }
                />
              </AltFormField>
              <AltFormField label={'1588 SYNC Message Interval'}>
                <OutlinedInput
                  data-testid="1588_sync_msg_interval"
                  placeholder="Enter 1588 SYNC Message Interval"
                  fullWidth={true}
                  value={optConfig.sync_1588_msg_interval}
                  onChange={({target}) =>
                    handleOptChange('sync_1588_msg_interval', target.value)
                  }
                />
              </AltFormField>
              <AltFormField label={'1588 SYNC Delay Request Message Interval'}>
                <OutlinedInput
                  data-testid="1588_sync_rq_msg_inverval"
                  placeholder="Enter 1588 SYNC Delay Request Message Interval"
                  fullWidth={true}
                  value={optConfig.sync_1588_delay_rq_msg_interval}
                  onChange={({target}) =>
                    handleOptChange(
                      'sync_1588_delay_rq_msg_interval',
                      target.value,
                    )
                  }
                />
              </AltFormField>
              <AltFormField label={'1588 SYNC Unicast/Multicast'}>
                <FormControl variant={'outlined'}>
                  <Select
                    value={optConfig.sync_1588_unicast_enable ? 1 : 0}
                    onChange={({target}) =>
                      handleOptChange(
                        'sync_1588_unicast_enable',
                        target.value === 1,
                      )
                    }
                    input={<OutlinedInput id="sync_1588_unicast_enable" />}>
                    <MenuItem value={0}>Multicast</MenuItem>
                    <MenuItem value={1}>Unicast</MenuItem>
                  </Select>
                </FormControl>
              </AltFormField>

              <AltFormField label={'1588 SYNC Unicast ServerIp'}>
                <OutlinedInput
                  data-testid="1588_sync_unicast_serverip"
                  placeholder="Enter 1588 SYNC Unicast ServerIp"
                  fullWidth={true}
                  value={optConfig.sync_1588_unicast_serverIp}
                  onChange={({target}) =>
                    handleOptChange('sync_1588_unicast_serverIp', target.value)
                  }
                />
              </AltFormField>
              <AltFormField label={'Transmit'}>
                <FormControl variant={'outlined'}>
                  <Select
                    value={config.transmit_enabled ? 1 : 0}
                    onChange={({target}) =>
                      handleEnbChange('transmit_enabled', target.value === 1)
                    }
                    input={<OutlinedInput id="transmitEnabled" />}>
                    <MenuItem value={0}>Disabled</MenuItem>
                    <MenuItem value={1}>Enabled</MenuItem>
                  </Select>
                </FormControl>
              </AltFormField>
            </>
          )}
        </List>
      </DialogContent>
      <DialogActions>
        <Button onClick={props.onClose} skin="regular">
          Cancel
        </Button>
        <Button onClick={onSave} variant="contained" color="primary">
          {props.isAdd ? 'Save And Add eNodeB' : 'Save'}
        </Button>
      </DialogActions>
    </>
  );
}

export function ConfigEdit(props: Props) {
  const [error, setError] = useState('');
  const {match} = useRouter();
  const enqueueSnackbar = useEnqueueSnackbar();
  const ctx = useContext(EnodebContext);
  const enodebSerial: string = match.params.enodebSerial;
  const enbInfo = ctx.state.enbInfo[enodebSerial];
  const [enb, setEnb] = useState<enodeb>(props.enb || DEFAULT_ENB_CONFIG);
  const onSave = async () => {
    try {
      if (props.isAdd) {
        // check if it is not a modify during add i.e we aren't switching tabs back
        // during add and modifying the information other than the serial number
        if (
          enb.serial in ctx.state.enbInfo &&
          enb.serial !== props.enb?.serial
        ) {
          setError(`eNodeB ${enb.serial} already exists`);
          return;
        }
      }

      if (enb.config == null) {
        enb.config = DEFAULT_ENB_CONFIG.config;
      }
      if (enb.enodeb_config == null || enb.enodeb_config.config_type == '') {
        enb.enodeb_config = {
          config_type: 'MANAGED',
          managed_config: DEFAULT_ENB_CONFIG.config,
        };
      }
      await ctx.setState(enb.serial, {
        enb_state: enbInfo?.enb_state ?? {},
        enb: enb,
      });
      if (props.enb) {
        enqueueSnackbar('eNodeb saved successfully', {
          variant: 'success',
        });
      }
      props.onSave(enb);
    } catch (e) {
      setError(e.response?.data?.message ?? e.message);
    }
  };

  return (
    <>
      <DialogContent data-testid="configEdit">
        <List>
          {error !== '' && (
            <AltFormField label={''}>
              <FormLabel data-testid="configEditError" error>
                {error}
              </FormLabel>
            </AltFormField>
          )}
          <AltFormField label={'Name'}>
            <OutlinedInput
              data-testid="name"
              placeholder="Enter Name"
              fullWidth={true}
              value={enb.name}
              onChange={({target}) => setEnb({...enb, name: target.value})}
            />
          </AltFormField>
          <AltFormField label={'Serial Number'}>
            <OutlinedInput
              data-testid="serial"
              placeholder="Ex: 12020000261814C0021"
              fullWidth={true}
              value={enb.serial}
              readOnly={props.enb ? true : false}
              onChange={({target}) => setEnb({...enb, serial: target.value})}
            />
          </AltFormField>
          <AltFormField label={'Description'}>
            <OutlinedInput
              data-testid="description"
              placeholder="Enter Description"
              fullWidth={true}
              multiline
              rows={4}
              value={enb.description}
              onChange={({target}) =>
                setEnb({...enb, description: target.value})
              }
            />
          </AltFormField>
        </List>
      </DialogContent>
      <DialogActions>
        <Button onClick={props.onClose} skin="regular">
          Cancel
        </Button>
        <Button onClick={onSave} variant="contained" color="primary">
          {props.isAdd ? 'Save And Continue' : 'Save'}
        </Button>
      </DialogActions>
    </>
  );
}

function coerceValue<T>(value: string, options: {[string]: T}): T {
  const values = Object.values(options);
  const keys = Object.keys(options);
  const optionKey = values.indexOf(value);
  if (optionKey > -1) {
    return options[keys[optionKey]];
  } else {
    throw Error('Expected a valid selection.');
  }
}

function isNumberInRange(value: string | number, lower: number, upper: number) {
  const val = parseInt(value, 10);
  if (isNaN(val)) {
    return false;
  }
  return val >= lower && val <= upper;
}

function buildRanConfig(config: enodeb_configuration, optConfig: OptConfig) {
  const response = {
    ...config,
    bandwidth_mhz: optConfig.bandwidthMhz,
    radioConfiguration: {
      powerControlParameters: {
        reference_signal_power: null,
        power_class: null,
        pb: null,
        pa: null,
      },
    },
    sync_1588: {
      sync_1588_switch: false,
      sync_1588_asymmetry: null,
      sync_1588_delay_rq_msg_interval: null,
      sync_1588_domain: null,
      sync_1588_holdover: null,
      sync_1588_msg_interval: null,
      sync_1588_unicast_enable: true,
      sync_1588_unicast_serverIp: null,
    },
    managementServer: {
      management_server_host: null,
      management_server_port: null,
      management_server_ssl_enable: false,
    },
  };
  if (!isNumberInRange(config.cell_id, 0, Math.pow(2, 28) - 1)) {
    throw Error('Invalid Configuration Cell ID. Valid range 0 - (2^28) - 1');
  }
  if (optConfig.earfcndl !== '') {
    if (!isNumberInRange(optConfig.earfcndl, 0, 65535)) {
      throw Error('Invalid EARFCNDL. Valid range 0 - 645535');
    }
    response['earfcndl'] = parseInt(optConfig.earfcndl);
  }

  if (optConfig.pci !== '') {
    if (!isNumberInRange(optConfig.pci, 0, 504)) {
      throw Error('Invalid PCI. Valid range 0 - 504');
    }
    response['pci'] = parseInt(optConfig.pci);
  }

  if (optConfig.specialSubframePattern !== '') {
    if (!isNumberInRange(optConfig.specialSubframePattern, 0, 9)) {
      throw Error('Invalid Special SubFrame Pattern, Valid range 0 - 9');
    }
    response['special_subframe_pattern'] = parseInt(
      optConfig.specialSubframePattern,
    );
  }

  if (optConfig.subframeAssignment !== '') {
    if (!isNumberInRange(optConfig.subframeAssignment, 0, 6)) {
      throw Error('Invalid SubFrame Assignment, Valid range 0 - 6');
    }
    response['subframe_assignment'] = parseInt(optConfig.subframeAssignment);
  }

  if (optConfig.tac !== '') {
    if (!isNumberInRange(optConfig.tac, 0, 65535)) {
      throw Error('Invalid TAC, Valid Range 0 - 65535');
    }
    response['tac'] = parseInt(optConfig.tac);
  }
  if (optConfig.reference_signal_power !== '') {
    response.radioConfiguration.powerControlParameters.reference_signal_power = parseInt(
      optConfig.reference_signal_power,
    );
  }
  if (optConfig.power_class !== '') {
    response.radioConfiguration.powerControlParameters.power_class = parseInt(
      optConfig.power_class,
    );
  }
  if (optConfig.pa !== '') {
    response.radioConfiguration.powerControlParameters.pa = parseInt(
      optConfig.pa,
    );
  }
  if (optConfig.pb !== '') {
    response.radioConfiguration.powerControlParameters.pb = parseInt(
      optConfig.pb,
    );
  }
  if (optConfig.mme_pool_1 !== '') {
    response['mme_pool_1'] = String(optConfig.mme_pool_1);
  }
  if (optConfig.mme_pool_2 !== '') {
    response['mme_pool_2'] = String(optConfig.mme_pool_2);
  }
  if (optConfig.management_server_ssl_enable !== '') {
    response.managementServer.management_server_ssl_enable = Boolean(
      optConfig.management_server_ssl_enable,
    );
  }
  if (optConfig.management_server_host !== '') {
    response.managementServer.management_server_host = String(
      optConfig.management_server_host,
    );
  }
  if (optConfig.management_server_port !== '') {
    response.managementServer.management_server_port = parseInt(
      optConfig.management_server_port,
    );
  }
  if (optConfig.sync_1588_switch !== '') {
    response.sync_1588.sync_1588_switch = Boolean(optConfig.sync_1588_switch);
  }

  if (optConfig.sync_1588_asymmetry !== '') {
    response.sync_1588.sync_1588_asymmetry = parseInt(
      optConfig.sync_1588_asymmetry,
    );
  }
  if (optConfig.sync_1588_delay_rq_msg_interval !== '') {
    response.sync_1588.sync_1588_delay_rq_msg_interval = parseInt(
      optConfig.sync_1588_delay_rq_msg_interval,
    );
  }
  if (optConfig.sync_1588_domain !== '') {
    response.sync_1588.sync_1588_domain = parseInt(optConfig.sync_1588_domain);
  }
  if (optConfig.sync_1588_holdover !== '') {
    response.sync_1588.sync_1588_holdover = parseInt(
      optConfig.sync_1588_holdover,
    );
  }
  if (optConfig.sync_1588_msg_interval !== '') {
    response.sync_1588.sync_1588_msg_interval = parseInt(
      optConfig.sync_1588_msg_interval,
    );
  }
  if (optConfig.sync_1588_unicast_serverIp !== '') {
    response.sync_1588.sync_1588_unicast_serverIp = String(
      optConfig.sync_1588_unicast_serverIp,
    );
  }
  if (optConfig.sync_1588_unicast_enable !== '') {
    response.sync_1588.sync_1588_unicast_enable = Boolean(
      optConfig.sync_1588_unicast_enable,
    );
  }
  return response;
}

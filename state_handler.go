// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package gonm

import (
	"fmt"
	"github.com/sofiworker/gonm/logger"
	"github.com/sofiworker/gonm/nm"
	"sync"

	dbus "github.com/godbus/dbus/v5"
	"github.com/sofiworker/gonm/dbusutil"
	nmdbus "github.com/sofiworker/gonm/system/org.freedesktop.networkmanager"
)

var vpnErrorTable = make(map[uint32]string)
var deviceErrorTable = make(map[uint32]string)

func initNmStateReasons() {
	// device error table
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_NONE] = "Device state changed"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_UNKNOWN] = "Device state changed, reason unknown"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_NOW_MANAGED] = "The device is now managed"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_NOW_UNMANAGED] = "The device is no longer managed"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_CONFIG_FAILED] = "The device has not been ready for configuration"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_IP_CONFIG_UNAVAILABLE] = "IP configuration could not be reserved (no available address, timeout, etc"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_IP_CONFIG_EXPIRED] = "The IP configuration is no longer valid"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_NO_SECRETS] = "Passwords were required but not provided"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_SUPPLICANT_DISCONNECT] = "The 802.1X supplicant disconnected from the access point or authentication server"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_SUPPLICANT_CONFIG_FAILED] = "Configuration of the 802.1X supplicant failed"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_SUPPLICANT_FAILED] = "The 802.1X supplicant quitted or failed unexpectedly"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_SUPPLICANT_TIMEOUT] = "The 802.1X supplicant took too long time to authenticate"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_PPP_START_FAILED] = "The PPP service failed to start within the allowed time"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_PPP_DISCONNECT] = "The PPP service disconnected unexpectedly"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_PPP_FAILED] = "The PPP service quitted or failed unexpectedly"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_DHCP_START_FAILED] = "The DHCP service failed to start within the allowed time"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_DHCP_ERROR] = "The DHCP service reported an unexpected error"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_DHCP_FAILED] = "The DHCP service quitted or failed unexpectedly"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_SHARED_START_FAILED] = "The shared connection service failed to start"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_SHARED_FAILED] = "The shared connection service quitted or failed unexpectedly"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_AUTOIP_START_FAILED] = "The AutoIP service failed to start"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_AUTOIP_ERROR] = "The AutoIP service reported an unexpected error"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_AUTOIP_FAILED] = "The AutoIP service quitted or failed unexpectedly"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_MODEM_BUSY] = "Dialing failed due to busy lines"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_MODEM_NO_DIAL_TONE] = "Dialing failed due to no dial tone"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_MODEM_NO_CARRIER] = "Dialing failed due to the carrier"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_MODEM_DIAL_TIMEOUT] = "Dialing timed out"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_MODEM_DIAL_FAILED] = "Dialing failed"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_MODEM_INIT_FAILED] = "Modem initialization failed"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_GSM_APN_FAILED] = "Failed to select the specified GSM APN"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_GSM_REGISTRATION_NOT_SEARCHING] = "No networks searched"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_GSM_REGISTRATION_DENIED] = "Network registration was denied"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_GSM_REGISTRATION_TIMEOUT] = "Network registration timed out"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_GSM_REGISTRATION_FAILED] = "Failed to register to the requested GSM network"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_GSM_PIN_CHECK_FAILED] = "PIN check failed"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_FIRMWARE_MISSING] = "Necessary firmware for the device may be missed"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_REMOVED] = "The device was removed"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_SLEEPING] = "NetworkManager went to sleep"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_CONNECTION_REMOVED] = "The device's active connection was removed or disappeared"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_USER_REQUESTED] = "A user or client requested to disconnect"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_CARRIER] = "The device's carrier/link changed" // TODO translate
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_CONNECTION_ASSUMED] = "The device's existing connection was assumed"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_SUPPLICANT_AVAILABLE] = "The 802.1x supplicant is now available" // TODO translate: full stop
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_MODEM_NOT_FOUND] = "The modem could not be found"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_BT_FAILED] = "The Bluetooth connection timed out or failed"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_GSM_SIM_NOT_INSERTED] = "GSM Modem's SIM Card was not inserted"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_GSM_SIM_PIN_REQUIRED] = "GSM Modem's SIM PIN required"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_GSM_SIM_PUK_REQUIRED] = "GSM Modem's SIM PUK required"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_GSM_SIM_WRONG] = "SIM card error in GSM Modem"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_INFINIBAND_MODE] = "InfiniBand device does not support connected mode"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_DEPENDENCY_FAILED] = "A dependency of the connection failed"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_BR2684_FAILED] = "RFC 2684 Ethernet bridging error to ADSL" // TODO translate
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_MODEM_MANAGER_UNAVAILABLE] = "ModemManager did not run or quitted unexpectedly"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_SSID_NOT_FOUND] = "The 802.11 WLAN network could not be found"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_SECONDARY_CONNECTION_FAILED] = "A secondary connection of the base connection failed"

	// works for nm 1.0+
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_DCB_FCOE_FAILED] = "DCB or FCoE setup failed"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_TEAMD_CONTROL_FAILED] = "Network teaming control failed"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_MODEM_FAILED] = "Modem failed to run or not available"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_MODEM_AVAILABLE] = "Modem now ready and available"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_SIM_PIN_INCORRECT] = "SIM PIN is incorrect"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_NEW_ACTIVATION] = "New connection activation is enqueuing"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_PARENT_CHANGED] = "Parent device changed"
	deviceErrorTable[nm.NM_DEVICE_STATE_REASON_PARENT_MANAGED_CHANGED] = "Management status of parent device changed"

	// device error table for custom state reasons
	deviceErrorTable[CUSTOM_NM_DEVICE_STATE_REASON_CABLE_UNPLUGGED] = "Network cable is unplugged"
	deviceErrorTable[CUSTOM_NM_DEVICE_STATE_REASON_MODEM_NO_SIGNAL] = "Please make sure SIM card has been inserted with mobile network signal"
	deviceErrorTable[CUSTOM_NM_DEVICE_STATE_REASON_MODEM_WRONG_PLAN] = "Please make sure a correct plan was selected without arrearage of SIM card"

	// vpn error table
	vpnErrorTable[nm.NM_VPN_CONNECTION_STATE_REASON_UNKNOWN] = "Failed to activate VPN connection, reason unknown"
	vpnErrorTable[nm.NM_VPN_CONNECTION_STATE_REASON_NONE] = "Failed to activate VPN connection"
	vpnErrorTable[nm.NM_VPN_CONNECTION_STATE_REASON_USER_DISCONNECTED] = "The VPN connection state changed due to being disconnected by users"
	vpnErrorTable[nm.NM_VPN_CONNECTION_STATE_REASON_DEVICE_DISCONNECTED] = "The VPN connection state changed due to being disconnected from devices"
	vpnErrorTable[nm.NM_VPN_CONNECTION_STATE_REASON_SERVICE_STOPPED] = "VPN service stopped"
	vpnErrorTable[nm.NM_VPN_CONNECTION_STATE_REASON_IP_CONFIG_INVALID] = "The IP config of VPN connection was invalid"
	vpnErrorTable[nm.NM_VPN_CONNECTION_STATE_REASON_CONNECT_TIMEOUT] = "The connection attempt to VPN service timed out"
	vpnErrorTable[nm.NM_VPN_CONNECTION_STATE_REASON_SERVICE_START_TIMEOUT] = "The VPN service start timed out"
	vpnErrorTable[nm.NM_VPN_CONNECTION_STATE_REASON_SERVICE_START_FAILED] = "The VPN service failed to start"
	vpnErrorTable[nm.NM_VPN_CONNECTION_STATE_REASON_NO_SECRETS] = "The VPN connection password was not provided"
	vpnErrorTable[nm.NM_VPN_CONNECTION_STATE_REASON_LOGIN_FAILED] = "Authentication to VPN server failed"
	vpnErrorTable[nm.NM_VPN_CONNECTION_STATE_REASON_CONNECTION_REMOVED] = "The connection was deleted from settings"
}

type stateHandler struct {
	m       *Manager
	devices map[dbus.ObjectPath]*deviceStateInfo
	locker  sync.Mutex

	sysSigLoop *dbusutil.SignalLoop
}

type deviceStateInfo struct {
	nmDev          nmdbus.Device
	enabled        bool
	devUdi         string
	devType        uint32
	aconnId        string
	connectionType string
}

func newStateHandler(sysSigLoop *dbusutil.SignalLoop, m *Manager) (sh *stateHandler) {
	sh = &stateHandler{
		m:          m,
		sysSigLoop: sysSigLoop,
		devices:    make(map[dbus.ObjectPath]*deviceStateInfo),
	}

	_, err := nmManager.ConnectDeviceRemoved(func(path dbus.ObjectPath) {
		sh.remove(path)
	})
	if err != nil {
		logger.Warn(err)
	}
	_, err = nmManager.ConnectDeviceAdded(func(path dbus.ObjectPath) {
		sh.watch(path)
	})
	if err != nil {
		logger.Warn(err)
	}
	for _, path := range nmGetDevices() {
		sh.watch(path)
	}

	err = nmManager.NetworkingEnabled().ConnectChanged(func(hasValue bool, value bool) {
		if !nmGetNetworkEnabled() {
			notifyAirplanModeEnabled()
		}
	})
	if err != nil {
		logger.Warn(err)
	}
	_ = nmManager.WirelessHardwareEnabled().ConnectChanged(func(hasValue bool, value bool) {
		if !nmGetWirelessHardwareEnabled() {
			notifyWirelessHardSwitchOff()
		}
	})

	return
}

func destroyStateHandler(sh *stateHandler) {
	for path := range sh.devices {
		sh.remove(path)
	}
	sh.devices = nil
}

func (sh *stateHandler) watch(path dbus.ObjectPath) {
	defer func() {
		if err := recover(); err != nil {
			sh.locker.Lock()
			defer sh.locker.Unlock()
			delete(sh.devices, path)
			logger.SError(err)
		}
	}()

	nmDev, err := nmNewDevice(path)
	if err != nil {
		return
	}

	// dont notify message when virtual device state changed
	if isVirtualDeviceIfc(nmDev) {
		return
	}

	deviceType, _ := nmDev.Device().DeviceType().Get(0)
	if !isDeviceTypeValid(deviceType) {
		return
	}

	sh.locker.Lock()
	defer sh.locker.Unlock()
	sh.devices[path] = &deviceStateInfo{nmDev: nmDev}
	sh.devices[path].devType = deviceType
	sh.devices[path].devUdi, _ = nmDev.Device().Udi().Get(0)
	enabled, err := sh.m.sysNetwork.IsDeviceEnabled(0, string(path))
	if err == nil {
		sh.devices[path].enabled = enabled
	} else {
		logger.Warn(err)
	}

	if data, err := nmGetDeviceActiveConnectionData(path); err == nil {
		// remember active connection id and type if exists
		sh.devices[path].aconnId = getSettingConnectionId(data)
		sh.devices[path].connectionType = getCustomConnectionType(data)
	}

	// connect signals
	nmDev.InitSignalExt(sh.sysSigLoop, true)
	_, err = nmDev.Device().ConnectStateChanged(func(newState, oldState, reason uint32) {
		var id string
		sh.m.activeConnectionsLock.Lock()
		for _, ac := range sh.m.activeConnections {
			// search dev
			for _, dev := range ac.Devices {
				// check if type is equal
				if dev == path {
					id = ac.Id
					break
				}
			}
		}
		sh.m.activeConnectionsLock.Unlock()
		logger.SDebugf("device state changed, %d => %d, reason[%d] %s", oldState, newState, reason, deviceErrorTable[reason])
		sh.locker.Lock()
		defer sh.locker.Unlock()
		// update id here
		if id != "" && id != "/" {
			sh.devices[path].aconnId = id
		}
		if data, err := nmGetDeviceActiveConnectionData(path); err == nil {
			// update active connection and type if exists
			sh.devices[path].connectionType = getCustomConnectionType(data)
		}
		dsi, ok := sh.devices[path]
		if !ok || (sh.devices[path].aconnId == "") {
			// the device already been removed
			return
		}

		switch newState {
		case nm.NM_DEVICE_STATE_PREPARE:
			if data, err := nmGetDeviceActiveConnectionData(path); err == nil {
				dsi.aconnId = getSettingConnectionId(data)
				icon := generalGetNotifyDisconnectedIcon(dsi.devType, path)
				logger.Debug("--------[Prepare] Active connection info:", dsi.aconnId, dsi.connectionType, dsi.nmDev.Path_())
				if dsi.connectionType == connectionWirelessHotspot {
					notify(icon, "", "Enabling hotspot")
				} else {
					// 防止连接状态由60变40再次弹出正在连接的通知消息
					if oldState == nm.NM_DEVICE_STATE_DISCONNECTED {
						notify(icon, "", fmt.Sprintf("Connecting %s", dsi.aconnId))
					}
				}
			}
		case nm.NM_DEVICE_STATE_ACTIVATED:
			icon := generalGetNotifyConnectedIcon(dsi.devType, path)
			msg := dsi.aconnId
			logger.Debug("--------[Activated] Active connection info:", dsi.aconnId, dsi.connectionType, dsi.nmDev.Path_())
			if dsi.connectionType == connectionWirelessHotspot {
				notify(icon, "", "Hotspot enabled")
			} else {
				notify(icon, "", fmt.Sprintf("%s connected", msg))
			}
		case nm.NM_DEVICE_STATE_FAILED, nm.NM_DEVICE_STATE_DISCONNECTED, nm.NM_DEVICE_STATE_NEED_AUTH,
			nm.NM_DEVICE_STATE_UNMANAGED, nm.NM_DEVICE_STATE_UNAVAILABLE:
			logger.SInfof("device disconnected, type %s, %d => %d, reason[%d] %s", getCustomDeviceType(dsi.devType), oldState, newState, reason, deviceErrorTable[reason])

			// ignore device removed signals for that could not
			// query related information correct
			if reason == nm.NM_DEVICE_STATE_REASON_REMOVED {
				if dsi.connectionType == connectionWirelessHotspot {
					icon := generalGetNotifyDisconnectedIcon(dsi.devType, path)
					notify(icon, "", "Hotspot disabled")
				}
				return
			}

			// ignore if device's old state is not available
			if !isDeviceStateAvailable(oldState) {
				logger.Debug("no notify, old state is not available")
				return
			}

			// notify only when network enabled
			if !nmGetNetworkEnabled() {
				logger.Debug("no notify, network disabled")
				return
			}

			// notify only when device enabled
			if oldState == nm.NM_DEVICE_STATE_DISCONNECTED && !dsi.enabled {
				logger.Debug("no notify, notify only when device enabled")
				return
			}

			// fix reasons
			switch dsi.devType {
			case nm.NM_DEVICE_TYPE_ETHERNET:
				if reason == nm.NM_DEVICE_STATE_REASON_CARRIER {
					reason = CUSTOM_NM_DEVICE_STATE_REASON_CABLE_UNPLUGGED
				}
			case nm.NM_DEVICE_TYPE_MODEM:
				if isDeviceStateReasonInvalid(reason) {
					// mobile device is specially, fix its reasons here
					signalQuality, _ := mmGetModemDeviceSignalQuality(dbus.ObjectPath(dsi.devUdi))
					if signalQuality == 0 {
						reason = CUSTOM_NM_DEVICE_STATE_REASON_MODEM_NO_SIGNAL
					} else {
						reason = CUSTOM_NM_DEVICE_STATE_REASON_MODEM_WRONG_PLAN
					}
				}
			}

			// ignore invalid reasons
			if isDeviceStateReasonInvalid(reason) {
				logger.Debug("no notify, device state reason invalid")
				return
			}

			logger.Debug("--------[Disconnect] Active connection info:", dsi.aconnId, dsi.connectionType, dsi.nmDev.Path_())
			var icon, msg string
			icon = generalGetNotifyDisconnectedIcon(dsi.devType, path)
			if len(msg) == 0 {
				switch reason {
				case nm.NM_DEVICE_STATE_REASON_NONE, nm.NM_DEVICE_STATE_REASON_USER_REQUESTED:
					if newState == nm.NM_DEVICE_STATE_DISCONNECTED || newState == nm.NM_DEVICE_STATE_UNAVAILABLE {
						if dsi.connectionType == connectionWirelessHotspot {
							notify(icon, "", "Hotspot disabled")
						} else {
							msg = fmt.Sprintf("%s disconnected", dsi.aconnId)
						}
					}
				case nm.NM_DEVICE_STATE_REASON_NEW_ACTIVATION:
				case nm.NM_DEVICE_STATE_REASON_IP_CONFIG_UNAVAILABLE:
					if dsi.connectionType == connectionWirelessHotspot {
						msg = "Unable to share hotspot, please check dnsmasq settings"
					} else if dsi.connectionType == connectionWireless {
						msg = fmt.Sprintf("Unable to connect %s, please keep closer to the wireless router", dsi.aconnId)
					} else if dsi.connectionType == connectionWired {
						msg = fmt.Sprintf("Unable to connect %s, please check your router or net cable.", dsi.aconnId)
					}
				case nm.NM_DEVICE_STATE_REASON_SUPPLICANT_DISCONNECT:
					if (oldState == nm.NM_DEVICE_STATE_CONFIG || oldState == nm.NM_DEVICE_STATE_ACTIVATED) && newState == nm.NM_DEVICE_STATE_NEED_AUTH {
						msg = fmt.Sprintf("Connection failed, unable to connect %s, wrong password", dsi.aconnId)
					} else if oldState == nm.NM_DEVICE_STATE_CONFIG && newState == nm.NM_DEVICE_STATE_FAILED {
						msg = fmt.Sprintf("Unable to connect %s",dsi.aconnId) 
					}
				case CUSTOM_NM_DEVICE_STATE_REASON_CABLE_UNPLUGGED: //disconnected due to cable unplugged
					// if device is ethernet,notify disconnected message

					logger.Debug("Disconnected due to unplugged cable")
					if dsi.devType == nm.NM_DEVICE_TYPE_ETHERNET {
						logger.Debug("unplugged device is ethernet")
						msg = fmt.Sprintf("%s disconnected", dsi.aconnId)
					}
				case nm.NM_DEVICE_STATE_REASON_NO_SECRETS:
					msg = fmt.Sprintf("Password is required to connect %s", dsi.aconnId)
				case nm.NM_DEVICE_STATE_REASON_SSID_NOT_FOUND:
					msg = fmt.Sprintf("The %s 802.11 WLAN network could not be found", dsi.aconnId)
					//default:
					//	if dsi.aconnId != "" {
					//		msg = fmt.Sprintf("%s disconnected"), dsi.aconnId)
					//	}
				}
			}
			if msg != "" {
				notify(icon, "", msg)
			}
		}
	})
	if err != nil {
		logger.Warn(err)
	}
}

func (sh *stateHandler) remove(path dbus.ObjectPath) {
	sh.locker.Lock()
	defer sh.locker.Unlock()
	if dev, ok := sh.devices[path]; ok {
		nmDestroyDevice(dev.nmDev)
		delete(sh.devices, path)
	}
}

// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package gonm

import (
	"github.com/sofiworker/gonm/logger"

	// libnotify "github.com/sofiworker/gonm/libs/notify"
	"github.com/sofiworker/gonm/libs/proxy"
	"github.com/sofiworker/gonm/loader"
	"github.com/sofiworker/gonm/proxychains"
)

var (
	manager *Manager
)

// func HandlePrepareForSleep(sleep bool) {
// 	if manager == nil {
// 		logger.Warning("Module 'network' has not start")
// 		return
// 	}
// 	if sleep {
// 		// suspend
// 		disableNotify()
// 		return
// 	}
// 	// wakeup
// 	enableNotify()
// 	//value decided the strategy of the wirelessScan
// 	_ = manager.RequestWirelessScan()
// 	time.AfterFunc(3*time.Second, func() {
// 		manager.clearAccessPoints()
// 	})
// }

type Module struct {
	*loader.ModuleBase
}

func NewModule() *Module {
	module := new(Module)
	module.ModuleBase = loader.NewModuleBase("network", module)
	return module
}

func (d *Module) GetDependencies() []string {
	return []string{}
}

func (d *Module) start() error {
	proxy.SetupProxy()

	service := loader.GetService()
	manager = NewManager(service)
	manager.init()

	managerServerObj, err := service.NewServerObject(dbusPath, manager, manager.syncConfig)
	if err != nil {
		return err
	}

	err = managerServerObj.SetWriteCallback(manager, "NetworkingEnabled", manager.networkingEnabledWriteCb)
	if err != nil {
		return err
	}
	err = managerServerObj.SetWriteCallback(manager, "VpnEnabled", manager.vpnEnabledWriteCb)
	if err != nil {
		return err
	}

	err = managerServerObj.Export()
	if err != nil {
		logger.SErrorf("failed to export manager:%v", err)
		manager = nil
		return err
	}

	manager.proxyChainsManager = proxychains.NewManager(service)
	err = service.Export(proxychains.DBusPath, manager.proxyChainsManager)
	if err != nil {
		logger.Warning("failed to export proxyChainsManager:", err)
		manager.proxyChainsManager = nil
		return err
	}

	err = service.RequestName(dbusServiceName)
	if err != nil {
		return err
	}

	err = manager.syncConfig.Register()
	if err != nil {
		logger.Warning("Failed to register sync service:", err)
	}

	initDBusDaemon()
	watchNetworkManagerRestart(manager)
	return nil
}

func (d *Module) Start() error {
	// libnotify.Init("dde-session-daemon")
	if manager != nil {
		return nil
	}

	initSlices() // initialize slice code
	initSysSignalLoop()
	// initNotifyManager()
	return d.start()
}

func (d *Module) Stop() error {
	if manager == nil {
		return nil
	}

	service := loader.GetService()

	err := service.ReleaseName(dbusServiceName)
	if err != nil {
		logger.Warn(err)
	}

	manager.destroy()
	destroyDBusDaemon()
	sysSigLoop.Stop()
	err = service.StopExport(manager)
	if err != nil {
		logger.Warn(err)
	}

	if manager.proxyChainsManager != nil {
		err = service.StopExport(manager.proxyChainsManager)
		if err != nil {
			logger.Warn(err)
		}
		manager.proxyChainsManager = nil
	}

	manager = nil
	return nil
}

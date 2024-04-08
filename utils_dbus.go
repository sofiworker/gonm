// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package gonm

import (
	"github.com/godbus/dbus/v5"
	"github.com/sofiworker/gonm/dbusutil"
	"github.com/sofiworker/gonm/dbusutil/proxy"
	"github.com/sofiworker/gonm/logger"
	dbusmgr "github.com/sofiworker/gonm/system/org.freedesktop.dbus"
	login1 "github.com/sofiworker/gonm/system/org.freedesktop.login1"
	nmdbus "github.com/sofiworker/gonm/system/org.freedesktop.networkmanager"
	notifications "github.com/sofiworker/gonm/system/session/org.freedesktop.notifications"
)

var (
	nmManager    nmdbus.Manager
	nmSettings   nmdbus.Settings
	loginManager login1.Manager
	dbusDaemon   dbusmgr.DBus // system dbus daemon
)

func (m *Manager) initDbusObjects() {
	systemBus, err := dbus.SystemBus()
	if err != nil {
		logger.SError(err)
		return
	}
	sessionBus, err := dbus.SessionBus()
	if err != nil {
		logger.SError(err)
		return
	}

	nmManager = nmdbus.NewManager(systemBus)
	nmManager.InitSignalExt(m.sysSigLoop, true)

	nmSettings = nmdbus.NewSettings(systemBus)
	nmSettings.InitSignalExt(m.sysSigLoop, true)

	loginManager = login1.NewManager(systemBus)
	loginManager.InitSignalExt(m.sysSigLoop, true)

	notification = notifications.NewNotifications(sessionBus)
}

var sysSigLoop *dbusutil.SignalLoop

func initSysSignalLoop() {
	systemBus, err := dbus.SystemBus()
	if err != nil {
		logger.SError(err)
		return
	}
	sysSigLoop = dbusutil.NewSignalLoop(systemBus, 50)
	sysSigLoop.Start()
}

func initDBusDaemon() {
	systemBus, err := dbus.SystemBus()
	if err != nil {
		logger.SError(err)
		return
	}
	dbusDaemon = dbusmgr.NewDBus(systemBus)
	dbusDaemon.InitSignalExt(sysSigLoop, true)
}

func destroyDBusDaemon() {
	dbusDaemon.RemoveHandler(proxy.RemoveAllHandlers)
}

func destroyDbusObjects() {
	// destroy global dbus objects manually when stopping service is
	// required for that there are multiple signal connected with
	// theme which need to be removed
	nmManager.RemoveHandler(proxy.RemoveAllHandlers)
	nmSettings.RemoveHandler(proxy.RemoveAllHandlers)
	loginManager.RemoveHandler(proxy.RemoveAllHandlers)
}

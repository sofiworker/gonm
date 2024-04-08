// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package proxychains

import (
	. "github.com/sofiworker/gonm/libs/gettext"
	"github.com/sofiworker/gonm/logger"
	notifications "github.com/sofiworker/gonm/system/session/org.freedesktop.notifications"
)

var (
	notification            notifications.Notifications
	notifyIconProxyEnabled  = "notification-network-proxy-enabled"
	notifyIconProxyDisabled = "notification-network-proxy-disabled"
)

func init() {
	sessionBus, err := dbus.SessionBus()
	if err != nil {
		notification = nil
		return
	}
	notification = notifications.NewNotifications(sessionBus)
}

func createNotify(appName string) func(string, string, string) {
	var nid uint32 = 0
	return func(icon, summary, body string) {
		if notification == nil {
			logger.Warning("notification is nil")
			logger.SDebugf("%s %s %s", icon, summary, body)
			return
		}
		var err error
		nid, err = notification.Notify(0, appName, nid,
			icon, summary, body, nil, nil, -1)
		if err != nil {
			logger.SWarn(err)
			return
		}
	}
}

var notify = createNotify("dde-control-center")

func notifyAppProxyEnabled() {
	notify(notifyIconProxyEnabled, Tr("Network"), Tr("Application proxy is set successfully"))
}
func notifyAppProxyEnableFailed() {
	notify(notifyIconProxyDisabled, Tr("Network"), Tr("Failed to set the application proxy"))
}

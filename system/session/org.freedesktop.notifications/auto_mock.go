// Code generated by "./generator ./session/org.freedesktop.notifications"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package notifications

import "fmt"
import "github.com/godbus/dbus/v5"
import "github.com/sofiworker/gonm/dbusutil"
import "github.com/sofiworker/gonm/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockNotifications struct {
	MockInterfaceNotifications // interface org.freedesktop.Notifications
	proxy.MockObject
}

type MockInterfaceNotifications struct {
	mock.Mock
}

// method GetCapabilities

func (v *MockInterfaceNotifications) GoGetCapabilities(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotifications) GetCapabilities(flags dbus.Flags) ([]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method Notify

func (v *MockInterfaceNotifications) GoNotify(flags dbus.Flags, ch chan *dbus.Call, app_name string, replaces_id uint32, app_icon string, summary string, body string, actions []string, hints map[string]dbus.Variant, expire_timeout int32) *dbus.Call {
	mockArgs := v.Called(flags, ch, app_name, replaces_id, app_icon, summary, body, actions, hints, expire_timeout)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotifications) Notify(flags dbus.Flags, app_name string, replaces_id uint32, app_icon string, summary string, body string, actions []string, hints map[string]dbus.Variant, expire_timeout int32) (uint32, error) {
	mockArgs := v.Called(flags, app_name, replaces_id, app_icon, summary, body, actions, hints, expire_timeout)

	ret0, ok := mockArgs.Get(0).(uint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method CloseNotification

func (v *MockInterfaceNotifications) GoCloseNotification(flags dbus.Flags, ch chan *dbus.Call, id uint32) *dbus.Call {
	mockArgs := v.Called(flags, ch, id)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotifications) CloseNotification(flags dbus.Flags, id uint32) error {
	mockArgs := v.Called(flags, id)

	return mockArgs.Error(0)
}

// method GetServerInformation

func (v *MockInterfaceNotifications) GoGetServerInformation(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceNotifications) GetServerInformation(flags dbus.Flags) (string, string, string, string, error) {
	mockArgs := v.Called(flags)

	return mockArgs.String(0), mockArgs.String(1), mockArgs.String(2), mockArgs.String(3), mockArgs.Error(4)
}

// signal NotificationClosed

func (v *MockInterfaceNotifications) ConnectNotificationClosed(cb func(id uint32, reason uint32)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// signal ActionInvoked

func (v *MockInterfaceNotifications) ConnectActionInvoked(cb func(id uint32, action_key string)) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

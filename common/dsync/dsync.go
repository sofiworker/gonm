// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package dsync

import (
	"encoding/json"
	"github.com/sofiworker/gonm/dbusutil"
	"github.com/sofiworker/gonm/dbusutil/proxy"
	"github.com/sofiworker/gonm/libs/strv"
	"github.com/sofiworker/gonm/logger"
	ofdbus "github.com/sofiworker/gonm/system/session/org.freedesktop.dbus"

	dbus "github.com/godbus/dbus/v5"
)

//go:generate dbusutil-gen em -type Config

type Interface interface {
	Get() (interface{}, error)
	Set(data []byte) error
}

type Config struct {
	name       string
	core       Interface
	dbusDaemon ofdbus.DBus
	path       dbus.ObjectPath
	sigLoop    *dbusutil.SignalLoop
}

const (
	// deepin-deepinid-daemon
	serviceName = "com.deepin.sync.Daemon"
	servicePath = "/com/deepin/sync/Daemon"
)

func NewConfig(name string, core Interface, sessionSigLoop *dbusutil.SignalLoop,
	path dbus.ObjectPath) *Config {
	c := &Config{
		name:    name,
		core:    core,
		sigLoop: sessionSigLoop,
		path:    path,
	}

	sessionBus := sessionSigLoop.Conn()
	c.dbusDaemon = ofdbus.NewDBus(sessionBus)
	c.dbusDaemon.InitSignalExt(sessionSigLoop, true)
	_, err := c.dbusDaemon.ConnectNameOwnerChanged(func(name string, oldOwner string, newOwner string) {
		if name == "com.deepin.sync.Daemon" && newOwner != "" {
			err := c.Register()
			if err != nil {
				logger.SError(err)
			}
		}
	})
	if err != nil {
		logger.SWarn(err)
	}
	return c
}

func (c *Config) Register() error {
	sessionBus, err := dbus.SessionBus()
	if err != nil {
		return err
	}

	availableServices, err := ofdbus.NewDBus(sessionBus).ListNames(0)
	if err != nil {
		return err
	}
	if !strv.Strv(availableServices).Contains(serviceName) {
		logger.Debug("sync daemon not exists")
		return nil
	}

	obj := sessionBus.Object(serviceName, servicePath)
	err = obj.Call("com.deepin.sync.Daemon.Register", 0, c.name, c.path).Err
	return err
}

func (c *Config) Destroy() {
	c.dbusDaemon.RemoveHandler(proxy.RemoveAllHandlers)
}

func (*Config) GetInterfaceName() string {
	return "com.deepin.sync.Config"
}

func (c *Config) Get() (data []byte, busErr *dbus.Error) {
	v, err := c.core.Get()
	if err != nil {
		return nil, dbusutil.ToError(err)
	}

	data, err = json.Marshal(v)
	if err != nil {
		return nil, dbusutil.ToError(err)
	}
	return data, nil
}

func (c *Config) Set(data []byte) *dbus.Error {
	err := c.core.Set(data)
	return dbusutil.ToError(err)
}

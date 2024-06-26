package gonm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"sync"

	"github.com/godbus/dbus/v5"
	"github.com/sofiworker/gonm/common/dsync"
	"github.com/sofiworker/gonm/dbusutil"
	"github.com/sofiworker/gonm/logger"
	"gopkg.in/ini.v1"
)

const (
	NmConnDir = "/etc/NetworkManager/system-connections"

	kfSectionConnection = "connection"
	kfSectionWIFI       = "wifi"
	kfKeyType           = "type"
	kfKeyMac            = "mac-address"
	kfKeyMacBlacklist   = "mac-address-blacklist"
	kfKeySeenBSSID      = "seen-bssids"
	kfKeyInterfaceName  = "interface-name"

	nmSyncVersion       = "1.0"
	nmService           = "org.freedesktop.NetworkManager"
	nmSettingsPath      = "/org/freedesktop/NetworkManager/Settings"
	nmSettingsIFC       = nmService + ".Settings"
	methodNMReloadConns = nmSettingsIFC + ".ReloadConnections"
)

var (
	instance      *GoNm
	gonmOnce      sync.Once
	nmSettingsObj dbus.BusObject
)

type Options struct {
	EnableNotification bool
	EnableProxy        bool
}

type GoNm struct {
	Options Options
}

func GetGoNm() *GoNm {
	gonmOnce.Do(func() {
		instance = &GoNm{}
	})
	return instance
}

func (g *GoNm) Start() error {
	module := NewModule()
	err := module.Start()
	if err != nil {
		logger.SErrorf("start network module failed:%v", err)
		return err
	}
	return nil
}

func NetworkGetConnections() (data []byte, busErr *dbus.Error) {
	list, err := getConnectionList(NmConnDir)
	if err != nil {
		return nil, dbusutil.ToError(err)
	}
	var info = dsync.NetworkData{
		Version:     nmSyncVersion,
		Connections: list,
	}
	data, err = json.Marshal(&info)
	if err != nil {
		return nil, dbusutil.ToError(err)
	}
	return data, nil
}

func NetworkSetConnections(data []byte) *dbus.Error {
	var info dsync.NetworkData
	err := json.Unmarshal(data, &info)
	if err != nil {
		return dbusutil.ToError(err)
	}
	err = info.Connections.Check()
	if err != nil {
		return dbusutil.ToError(err)
	}
	list, _ := getConnectionList(NmConnDir)
	for _, conn := range info.Connections {
		tmp := list.Get(conn.Type, conn.Filename)
		if tmp != nil && tmp.Equal(conn) {
			continue
		}
		// add or modify
		err = conn.WriteFile(NmConnDir)
		if err != nil {
			// TODO(jouyouyun): handle error
			logger.SErrorf("[Network] Failed to write connection file:%v",
				err)
			// return dbusutil.ToError(err)
			continue
		}
	}
	removeList := info.Connections.Diff(list)
	for _, conn := range removeList {
		err = conn.RemoveFile(NmConnDir)
		if err != nil {
			// TODO(jouyouyun): handle error
			logger.SErrorf("[Network] Failed to remove connection file:%v", err)
			continue
		}
	}
	err = reloadConnections()
	if err != nil {
		logger.Warning("[Network] Failed to reload connections:", err)
	}
	return nil
}

func getConnectionList(dir string) (dsync.ConnectionList, error) {
	files, err := getConnectionFiles(dir)
	if err != nil {
		return nil, err
	}

	var datas dsync.ConnectionList
	for _, f := range files {
		data, err := loadConnectionFile(f)
		if err != nil {
			continue
		}
		if datas.Exists(data) {
			continue
		}
		datas = append(datas, data)
	}
	sort.Sort(datas)
	return datas, nil
}

func loadConnectionFile(filename string) (*dsync.Connection, error) {
	cfg, err := ini.Load(filename)
	ty := cfg.Section(kfSectionConnection).Key(kfKeyType).String()
	if ty != dsync.ConnTypeWIFI {
		return nil, dsync.ErrConnUnsupportedType
	}

	cfg.Section(kfSectionConnection).DeleteKey(kfKeyInterfaceName)
	cfg.Section(kfSectionWIFI).DeleteKey(kfKeyMac)
	cfg.Section(kfSectionWIFI).DeleteKey(kfKeyMacBlacklist)
	cfg.Section(kfSectionWIFI).DeleteKey(kfKeySeenBSSID)

	var bs bytes.Buffer
	_, err = cfg.WriteTo(&bs)
	if err != nil {
		return nil, err
	}

	return &dsync.Connection{
		Type:     ty,
		Filename: filepath.Base(filename),
		Contents: []byte(bs.Bytes()),
	}, nil
}

func getConnectionFiles(dir string) ([]string, error) {
	finfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var files []string
	for _, finfo := range finfos {
		if finfo.IsDir() {
			continue
		}
		files = append(files, filepath.Join(dir, finfo.Name()))
	}
	return files, nil
}

func reloadConnections() error {
	obj, err := newSettingsBus()
	if err != nil {
		return err
	}
	var success bool
	err = obj.Call(methodNMReloadConns, 0).Store(&success)
	if err != nil {
		return err
	}
	if !success {
		return fmt.Errorf("reload connections failed")
	}
	return nil
}

func newSettingsBus() (dbus.BusObject, error) {
	if nmSettingsObj != nil {
		return nmSettingsObj, nil
	}
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}
	nmSettingsObj = conn.Object(nmService, nmSettingsPath)
	return nmSettingsObj, nil
}

func startBacklightHelperAsync(conn *dbus.Conn) {
	go func() {
		obj := conn.Object("org.deepin.dde.BacklightHelper1", "/org/deepin/dde/BacklightHelper1")
		err := obj.Call("org.freedesktop.DBus.Peer.Ping", 0).Err

		if err != nil {
			logger.Warn(err)
		}
	}()
}

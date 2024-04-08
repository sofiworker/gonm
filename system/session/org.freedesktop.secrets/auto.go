// Code generated by "./generator ./session/org.freedesktop.secrets"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package secrets

import "errors"
import "fmt"
import "github.com/godbus/dbus/v5"

import "github.com/sofiworker/gonm/dbusutil"
import "github.com/sofiworker/gonm/dbusutil/proxy"
import "unsafe"

type Service interface {
	service // interface org.freedesktop.Secret.Service
	proxy.Object
}

type objectService struct {
	interfaceService // interface org.freedesktop.Secret.Service
	proxy.ImplObject
}

func NewService(conn *dbus.Conn) Service {
	obj := new(objectService)
	obj.ImplObject.Init_(conn, "org.freedesktop.secrets", "/org/freedesktop/secrets")
	return obj
}

type service interface {
	GoOpenSession(flags dbus.Flags, ch chan *dbus.Call, algorithm string, input dbus.Variant) *dbus.Call
	OpenSession(flags dbus.Flags, algorithm string, input dbus.Variant) (dbus.Variant, dbus.ObjectPath, error)
	GoCreateCollection(flags dbus.Flags, ch chan *dbus.Call, properties map[string]dbus.Variant, alias string) *dbus.Call
	CreateCollection(flags dbus.Flags, properties map[string]dbus.Variant, alias string) (dbus.ObjectPath, dbus.ObjectPath, error)
	GoSearchItems(flags dbus.Flags, ch chan *dbus.Call, attributes map[string]string) *dbus.Call
	SearchItems(flags dbus.Flags, attributes map[string]string) ([]dbus.ObjectPath, []dbus.ObjectPath, error)
	GoUnlock(flags dbus.Flags, ch chan *dbus.Call, objects []dbus.ObjectPath) *dbus.Call
	Unlock(flags dbus.Flags, objects []dbus.ObjectPath) ([]dbus.ObjectPath, dbus.ObjectPath, error)
	GoLock(flags dbus.Flags, ch chan *dbus.Call, objects []dbus.ObjectPath) *dbus.Call
	Lock(flags dbus.Flags, objects []dbus.ObjectPath) ([]dbus.ObjectPath, dbus.ObjectPath, error)
	GoLockService(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	LockService(flags dbus.Flags) error
	GoChangeLock(flags dbus.Flags, ch chan *dbus.Call, collection dbus.ObjectPath) *dbus.Call
	ChangeLock(flags dbus.Flags, collection dbus.ObjectPath) (dbus.ObjectPath, error)
	GoGetSecrets(flags dbus.Flags, ch chan *dbus.Call, items []dbus.ObjectPath, session dbus.ObjectPath) *dbus.Call
	GetSecrets(flags dbus.Flags, items []dbus.ObjectPath, session dbus.ObjectPath) (map[dbus.ObjectPath]Secret, error)
	GoReadAlias(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call
	ReadAlias(flags dbus.Flags, name string) (dbus.ObjectPath, error)
	GoSetAlias(flags dbus.Flags, ch chan *dbus.Call, name string, collection dbus.ObjectPath) *dbus.Call
	SetAlias(flags dbus.Flags, name string, collection dbus.ObjectPath) error
	ConnectCollectionCreated(cb func(collection dbus.ObjectPath)) (dbusutil.SignalHandlerId, error)
	ConnectCollectionDeleted(cb func(collection dbus.ObjectPath)) (dbusutil.SignalHandlerId, error)
	ConnectCollectionChanged(cb func(collection dbus.ObjectPath)) (dbusutil.SignalHandlerId, error)
	Collections() proxy.PropObjectPathArray
}

type interfaceService struct{}

func (v *interfaceService) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceService) GetInterfaceName_() string {
	return "org.freedesktop.Secret.Service"
}

// method OpenSession

func (v *interfaceService) GoOpenSession(flags dbus.Flags, ch chan *dbus.Call, algorithm string, input dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".OpenSession", flags, ch, algorithm, input)
}

func (*interfaceService) StoreOpenSession(call *dbus.Call) (output dbus.Variant, result dbus.ObjectPath, err error) {
	err = call.Store(&output, &result)
	return
}

func (v *interfaceService) OpenSession(flags dbus.Flags, algorithm string, input dbus.Variant) (dbus.Variant, dbus.ObjectPath, error) {
	return v.StoreOpenSession(
		<-v.GoOpenSession(flags, make(chan *dbus.Call, 1), algorithm, input).Done)
}

// method CreateCollection

func (v *interfaceService) GoCreateCollection(flags dbus.Flags, ch chan *dbus.Call, properties map[string]dbus.Variant, alias string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateCollection", flags, ch, properties, alias)
}

func (*interfaceService) StoreCreateCollection(call *dbus.Call) (collection dbus.ObjectPath, prompt dbus.ObjectPath, err error) {
	err = call.Store(&collection, &prompt)
	return
}

func (v *interfaceService) CreateCollection(flags dbus.Flags, properties map[string]dbus.Variant, alias string) (dbus.ObjectPath, dbus.ObjectPath, error) {
	return v.StoreCreateCollection(
		<-v.GoCreateCollection(flags, make(chan *dbus.Call, 1), properties, alias).Done)
}

// method SearchItems

func (v *interfaceService) GoSearchItems(flags dbus.Flags, ch chan *dbus.Call, attributes map[string]string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SearchItems", flags, ch, attributes)
}

func (*interfaceService) StoreSearchItems(call *dbus.Call) (unlocked []dbus.ObjectPath, locked []dbus.ObjectPath, err error) {
	err = call.Store(&unlocked, &locked)
	return
}

func (v *interfaceService) SearchItems(flags dbus.Flags, attributes map[string]string) ([]dbus.ObjectPath, []dbus.ObjectPath, error) {
	return v.StoreSearchItems(
		<-v.GoSearchItems(flags, make(chan *dbus.Call, 1), attributes).Done)
}

// method Unlock

func (v *interfaceService) GoUnlock(flags dbus.Flags, ch chan *dbus.Call, objects []dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Unlock", flags, ch, objects)
}

func (*interfaceService) StoreUnlock(call *dbus.Call) (unlocked []dbus.ObjectPath, prompt dbus.ObjectPath, err error) {
	err = call.Store(&unlocked, &prompt)
	return
}

func (v *interfaceService) Unlock(flags dbus.Flags, objects []dbus.ObjectPath) ([]dbus.ObjectPath, dbus.ObjectPath, error) {
	return v.StoreUnlock(
		<-v.GoUnlock(flags, make(chan *dbus.Call, 1), objects).Done)
}

// method Lock

func (v *interfaceService) GoLock(flags dbus.Flags, ch chan *dbus.Call, objects []dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Lock", flags, ch, objects)
}

func (*interfaceService) StoreLock(call *dbus.Call) (locked []dbus.ObjectPath, Prompt dbus.ObjectPath, err error) {
	err = call.Store(&locked, &Prompt)
	return
}

func (v *interfaceService) Lock(flags dbus.Flags, objects []dbus.ObjectPath) ([]dbus.ObjectPath, dbus.ObjectPath, error) {
	return v.StoreLock(
		<-v.GoLock(flags, make(chan *dbus.Call, 1), objects).Done)
}

// method LockService

func (v *interfaceService) GoLockService(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LockService", flags, ch)
}

func (v *interfaceService) LockService(flags dbus.Flags) error {
	return (<-v.GoLockService(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method ChangeLock

func (v *interfaceService) GoChangeLock(flags dbus.Flags, ch chan *dbus.Call, collection dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ChangeLock", flags, ch, collection)
}

func (*interfaceService) StoreChangeLock(call *dbus.Call) (prompt dbus.ObjectPath, err error) {
	err = call.Store(&prompt)
	return
}

func (v *interfaceService) ChangeLock(flags dbus.Flags, collection dbus.ObjectPath) (dbus.ObjectPath, error) {
	return v.StoreChangeLock(
		<-v.GoChangeLock(flags, make(chan *dbus.Call, 1), collection).Done)
}

// method GetSecrets

func (v *interfaceService) GoGetSecrets(flags dbus.Flags, ch chan *dbus.Call, items []dbus.ObjectPath, session dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSecrets", flags, ch, items, session)
}

func (*interfaceService) StoreGetSecrets(call *dbus.Call) (secrets map[dbus.ObjectPath]Secret, err error) {
	err = call.Store(&secrets)
	return
}

func (v *interfaceService) GetSecrets(flags dbus.Flags, items []dbus.ObjectPath, session dbus.ObjectPath) (map[dbus.ObjectPath]Secret, error) {
	return v.StoreGetSecrets(
		<-v.GoGetSecrets(flags, make(chan *dbus.Call, 1), items, session).Done)
}

// method ReadAlias

func (v *interfaceService) GoReadAlias(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReadAlias", flags, ch, name)
}

func (*interfaceService) StoreReadAlias(call *dbus.Call) (collection dbus.ObjectPath, err error) {
	err = call.Store(&collection)
	return
}

func (v *interfaceService) ReadAlias(flags dbus.Flags, name string) (dbus.ObjectPath, error) {
	return v.StoreReadAlias(
		<-v.GoReadAlias(flags, make(chan *dbus.Call, 1), name).Done)
}

// method SetAlias

func (v *interfaceService) GoSetAlias(flags dbus.Flags, ch chan *dbus.Call, name string, collection dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAlias", flags, ch, name, collection)
}

func (v *interfaceService) SetAlias(flags dbus.Flags, name string, collection dbus.ObjectPath) error {
	return (<-v.GoSetAlias(flags, make(chan *dbus.Call, 1), name, collection).Done).Err
}

// signal CollectionCreated

func (v *interfaceService) ConnectCollectionCreated(cb func(collection dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "CollectionCreated", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".CollectionCreated",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var collection dbus.ObjectPath
		err := dbus.Store(sig.Body, &collection)
		if err == nil {
			cb(collection)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal CollectionDeleted

func (v *interfaceService) ConnectCollectionDeleted(cb func(collection dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "CollectionDeleted", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".CollectionDeleted",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var collection dbus.ObjectPath
		err := dbus.Store(sig.Body, &collection)
		if err == nil {
			cb(collection)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal CollectionChanged

func (v *interfaceService) ConnectCollectionChanged(cb func(collection dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "CollectionChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".CollectionChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var collection dbus.ObjectPath
		err := dbus.Store(sig.Body, &collection)
		if err == nil {
			cb(collection)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Collections ao

func (v *interfaceService) Collections() proxy.PropObjectPathArray {
	return &proxy.ImplPropObjectPathArray{
		Impl: v,
		Name: "Collections",
	}
}

type Collection interface {
	collection // interface org.freedesktop.Secret.Collection
	proxy.Object
}

type objectCollection struct {
	interfaceCollection // interface org.freedesktop.Secret.Collection
	proxy.ImplObject
}

func NewCollection(conn *dbus.Conn, path dbus.ObjectPath) (Collection, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectCollection)
	obj.ImplObject.Init_(conn, "org.freedesktop.secrets", path)
	return obj, nil
}

type collection interface {
	GoDelete(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Delete(flags dbus.Flags) (dbus.ObjectPath, error)
	GoSearchItems(flags dbus.Flags, ch chan *dbus.Call, attributes map[string]string) *dbus.Call
	SearchItems(flags dbus.Flags, attributes map[string]string) ([]dbus.ObjectPath, error)
	GoCreateItem(flags dbus.Flags, ch chan *dbus.Call, properties map[string]dbus.Variant, secret Secret, replace bool) *dbus.Call
	CreateItem(flags dbus.Flags, properties map[string]dbus.Variant, secret Secret, replace bool) (dbus.ObjectPath, dbus.ObjectPath, error)
	ConnectItemCreated(cb func(item dbus.ObjectPath)) (dbusutil.SignalHandlerId, error)
	ConnectItemDeleted(cb func(item dbus.ObjectPath)) (dbusutil.SignalHandlerId, error)
	ConnectItemChanged(cb func(item dbus.ObjectPath)) (dbusutil.SignalHandlerId, error)
	Items() proxy.PropObjectPathArray
	Label() proxy.PropString
	Locked() proxy.PropBool
	Created() proxy.PropUint64
	Modified() proxy.PropUint64
}

type interfaceCollection struct{}

func (v *interfaceCollection) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceCollection) GetInterfaceName_() string {
	return "org.freedesktop.Secret.Collection"
}

// method Delete

func (v *interfaceCollection) GoDelete(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Delete", flags, ch)
}

func (*interfaceCollection) StoreDelete(call *dbus.Call) (prompt dbus.ObjectPath, err error) {
	err = call.Store(&prompt)
	return
}

func (v *interfaceCollection) Delete(flags dbus.Flags) (dbus.ObjectPath, error) {
	return v.StoreDelete(
		<-v.GoDelete(flags, make(chan *dbus.Call, 1)).Done)
}

// method SearchItems

func (v *interfaceCollection) GoSearchItems(flags dbus.Flags, ch chan *dbus.Call, attributes map[string]string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SearchItems", flags, ch, attributes)
}

func (*interfaceCollection) StoreSearchItems(call *dbus.Call) (results []dbus.ObjectPath, err error) {
	err = call.Store(&results)
	return
}

func (v *interfaceCollection) SearchItems(flags dbus.Flags, attributes map[string]string) ([]dbus.ObjectPath, error) {
	return v.StoreSearchItems(
		<-v.GoSearchItems(flags, make(chan *dbus.Call, 1), attributes).Done)
}

// method CreateItem

func (v *interfaceCollection) GoCreateItem(flags dbus.Flags, ch chan *dbus.Call, properties map[string]dbus.Variant, secret Secret, replace bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreateItem", flags, ch, properties, secret, replace)
}

func (*interfaceCollection) StoreCreateItem(call *dbus.Call) (item dbus.ObjectPath, prompt dbus.ObjectPath, err error) {
	err = call.Store(&item, &prompt)
	return
}

func (v *interfaceCollection) CreateItem(flags dbus.Flags, properties map[string]dbus.Variant, secret Secret, replace bool) (dbus.ObjectPath, dbus.ObjectPath, error) {
	return v.StoreCreateItem(
		<-v.GoCreateItem(flags, make(chan *dbus.Call, 1), properties, secret, replace).Done)
}

// signal ItemCreated

func (v *interfaceCollection) ConnectItemCreated(cb func(item dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ItemCreated", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ItemCreated",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var item dbus.ObjectPath
		err := dbus.Store(sig.Body, &item)
		if err == nil {
			cb(item)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ItemDeleted

func (v *interfaceCollection) ConnectItemDeleted(cb func(item dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ItemDeleted", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ItemDeleted",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var item dbus.ObjectPath
		err := dbus.Store(sig.Body, &item)
		if err == nil {
			cb(item)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ItemChanged

func (v *interfaceCollection) ConnectItemChanged(cb func(item dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ItemChanged", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ItemChanged",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var item dbus.ObjectPath
		err := dbus.Store(sig.Body, &item)
		if err == nil {
			cb(item)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Items ao

func (v *interfaceCollection) Items() proxy.PropObjectPathArray {
	return &proxy.ImplPropObjectPathArray{
		Impl: v,
		Name: "Items",
	}
}

// property Label s

func (v *interfaceCollection) Label() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Label",
	}
}

// property Locked b

func (v *interfaceCollection) Locked() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "Locked",
	}
}

// property Created t

func (v *interfaceCollection) Created() proxy.PropUint64 {
	return &proxy.ImplPropUint64{
		Impl: v,
		Name: "Created",
	}
}

// property Modified t

func (v *interfaceCollection) Modified() proxy.PropUint64 {
	return &proxy.ImplPropUint64{
		Impl: v,
		Name: "Modified",
	}
}

type Item interface {
	item // interface org.freedesktop.Secret.Item
	proxy.Object
}

type objectItem struct {
	interfaceItem // interface org.freedesktop.Secret.Item
	proxy.ImplObject
}

func NewItem(conn *dbus.Conn, path dbus.ObjectPath) (Item, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectItem)
	obj.ImplObject.Init_(conn, "org.freedesktop.secrets", path)
	return obj, nil
}

type item interface {
	GoDelete(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Delete(flags dbus.Flags) (dbus.ObjectPath, error)
	GoGetSecret(flags dbus.Flags, ch chan *dbus.Call, session dbus.ObjectPath) *dbus.Call
	GetSecret(flags dbus.Flags, session dbus.ObjectPath) (Secret, error)
	GoSetSecret(flags dbus.Flags, ch chan *dbus.Call, secret Secret) *dbus.Call
	SetSecret(flags dbus.Flags, secret Secret) error
	Locked() proxy.PropBool
	Attributes() PropItemAttributes
	Label() proxy.PropString
	Type() proxy.PropString
	Created() proxy.PropUint64
	Modified() proxy.PropUint64
}

type interfaceItem struct{}

func (v *interfaceItem) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceItem) GetInterfaceName_() string {
	return "org.freedesktop.Secret.Item"
}

// method Delete

func (v *interfaceItem) GoDelete(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Delete", flags, ch)
}

func (*interfaceItem) StoreDelete(call *dbus.Call) (Prompt dbus.ObjectPath, err error) {
	err = call.Store(&Prompt)
	return
}

func (v *interfaceItem) Delete(flags dbus.Flags) (dbus.ObjectPath, error) {
	return v.StoreDelete(
		<-v.GoDelete(flags, make(chan *dbus.Call, 1)).Done)
}

// method GetSecret

func (v *interfaceItem) GoGetSecret(flags dbus.Flags, ch chan *dbus.Call, session dbus.ObjectPath) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSecret", flags, ch, session)
}

func (*interfaceItem) StoreGetSecret(call *dbus.Call) (secret Secret, err error) {
	err = call.Store(&secret)
	return
}

func (v *interfaceItem) GetSecret(flags dbus.Flags, session dbus.ObjectPath) (Secret, error) {
	return v.StoreGetSecret(
		<-v.GoGetSecret(flags, make(chan *dbus.Call, 1), session).Done)
}

// method SetSecret

func (v *interfaceItem) GoSetSecret(flags dbus.Flags, ch chan *dbus.Call, secret Secret) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetSecret", flags, ch, secret)
}

func (v *interfaceItem) SetSecret(flags dbus.Flags, secret Secret) error {
	return (<-v.GoSetSecret(flags, make(chan *dbus.Call, 1), secret).Done).Err
}

// property Locked b

func (v *interfaceItem) Locked() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "Locked",
	}
}

type PropItemAttributes interface {
	Get(flags dbus.Flags) (value map[string]string, err error)
	Set(flags dbus.Flags, value map[string]string) error
	ConnectChanged(cb func(hasValue bool, value map[string]string)) error
}

type implPropItemAttributes struct {
	Impl proxy.Implementer
	Name string
}

func (p implPropItemAttributes) Get(flags dbus.Flags) (value map[string]string, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p implPropItemAttributes) Set(flags dbus.Flags, value map[string]string) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p implPropItemAttributes) ConnectChanged(cb func(hasValue bool, value map[string]string)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v map[string]string
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}

// property Attributes a{ss}

func (v *interfaceItem) Attributes() PropItemAttributes {
	return &implPropItemAttributes{
		Impl: v,
		Name: "Attributes",
	}
}

// property Label s

func (v *interfaceItem) Label() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Label",
	}
}

// property Type s

func (v *interfaceItem) Type() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Type",
	}
}

// property Created t

func (v *interfaceItem) Created() proxy.PropUint64 {
	return &proxy.ImplPropUint64{
		Impl: v,
		Name: "Created",
	}
}

// property Modified t

func (v *interfaceItem) Modified() proxy.PropUint64 {
	return &proxy.ImplPropUint64{
		Impl: v,
		Name: "Modified",
	}
}

type Session interface {
	session // interface org.freedesktop.Secret.Session
	proxy.Object
}

type objectSession struct {
	interfaceSession // interface org.freedesktop.Secret.Session
	proxy.ImplObject
}

func NewSession(conn *dbus.Conn, path dbus.ObjectPath) (Session, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectSession)
	obj.ImplObject.Init_(conn, "org.freedesktop.secrets", path)
	return obj, nil
}

type session interface {
	GoClose(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Close(flags dbus.Flags) error
}

type interfaceSession struct{}

func (v *interfaceSession) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceSession) GetInterfaceName_() string {
	return "org.freedesktop.Secret.Session"
}

// method Close

func (v *interfaceSession) GoClose(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Close", flags, ch)
}

func (v *interfaceSession) Close(flags dbus.Flags) error {
	return (<-v.GoClose(flags, make(chan *dbus.Call, 1)).Done).Err
}

type Prompt interface {
	prompt // interface org.freedesktop.Secret.Prompt
	proxy.Object
}

type objectPrompt struct {
	interfacePrompt // interface org.freedesktop.Secret.Prompt
	proxy.ImplObject
}

func NewPrompt(conn *dbus.Conn, path dbus.ObjectPath) (Prompt, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(objectPrompt)
	obj.ImplObject.Init_(conn, "org.freedesktop.secrets", path)
	return obj, nil
}

type prompt interface {
	GoPrompt(flags dbus.Flags, ch chan *dbus.Call, window_id string) *dbus.Call
	Prompt(flags dbus.Flags, window_id string) error
	GoDismiss(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Dismiss(flags dbus.Flags) error
	ConnectCompleted(cb func(dismissed bool, result dbus.Variant)) (dbusutil.SignalHandlerId, error)
}

type interfacePrompt struct{}

func (v *interfacePrompt) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfacePrompt) GetInterfaceName_() string {
	return "org.freedesktop.Secret.Prompt"
}

// method Prompt

func (v *interfacePrompt) GoPrompt(flags dbus.Flags, ch chan *dbus.Call, window_id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Prompt", flags, ch, window_id)
}

func (v *interfacePrompt) Prompt(flags dbus.Flags, window_id string) error {
	return (<-v.GoPrompt(flags, make(chan *dbus.Call, 1), window_id).Done).Err
}

// method Dismiss

func (v *interfacePrompt) GoDismiss(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Dismiss", flags, ch)
}

func (v *interfacePrompt) Dismiss(flags dbus.Flags) error {
	return (<-v.GoDismiss(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal Completed

func (v *interfacePrompt) ConnectCompleted(cb func(dismissed bool, result dbus.Variant)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "Completed", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".Completed",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var dismissed bool
		var result dbus.Variant
		err := dbus.Store(sig.Body, &dismissed, &result)
		if err == nil {
			cb(dismissed, result)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

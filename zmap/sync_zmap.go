package zmap

import (
	"sync"

	"gitee.com/sienectagv/gozk/zutils"
)

// type (
// 	Key = zutils.Key
// 	Val = zutils.Val
// )

type SyncMap sync.Map

func NewSyncMap() *SyncMap {
	return &SyncMap{}
}

func FromStringMap(in Map) *SyncMap {
	m := NewSyncMap()
	for k, v := range in {
		m.Insert(k, v)
	}
	return m
}

func (sm *SyncMap) InsertMap(in Map) {
	for k, v := range in {
		sm.Insert(k, v)
	}
}

func (sm *SyncMap) InsertMapByKeys(in Map, keys ...Key) {
	if nil == in {
		return
	}
	for _, k := range keys {
		if v, ok := in[k]; ok {
			sm.Insert(k, v)
		}
	}
}

func (sm *SyncMap) CopyTo(out Map) {
	(*sync.Map)(sm).Range(func(k Key, v Val) bool {
		out.Insert(k, v)
		return true
	})
}

func (sm *SyncMap) CopyToByKeys(out Map, keys ...Key) {
	for _, k := range keys {
		if v, ok := sm.Get(k); ok {
			out.Insert(k, v)
		}
	}
}

func (sm *SyncMap) Get(key Key) (Val, bool) {
	return (*sync.Map)(sm).Load(key)
}

func (sm *SyncMap) Insert(k Key, val Val) {
	(*sync.Map)(sm).Store(k, val)
}

func (sm *SyncMap) Delete(key Key) (Val, bool) {
	val, ok := (*sync.Map)(sm).Load(key)
	sm.Delete(key)
	return val, ok
}

func (sm *SyncMap) Value(k Key) (Val, bool) {
	return nil, false
}

func (sm *SyncMap) Len() int {
	l := 0
	(*sync.Map)(sm).Range(func(k Key, v Val) bool {
		l += 1
		return true
	})
	return l
}

func (sm *SyncMap) GetString(k string) (string, bool) {
	v, ok := sm.Value(k)
	if !ok {
		return "", ok
	}
	return zutils.InterfaceToString(v)
}

func (sm *SyncMap) GetInt64(k string) (int64, bool) {
	v, ok := sm.Value(k)
	if !ok {
		return 0, ok
	}
	if i, ok := zutils.InterfaceToInt(v); ok {
		return i, ok
	}
	if u, ok := zutils.InterfaceToUint(v); ok {
		return int64(u), ok
	}
	return 0, false
}

func (sm *SyncMap) GetInt(k string) (int, bool) {
	if v, ok := sm.GetInt64(k); ok {
		return int(v), ok
	}
	return 0, false
}

func (sm *SyncMap) TryInt(k string, v *int) bool {
	i, ok := sm.GetInt(k)
	if ok {
		*v = i
	}
	return ok
}

func (sm *SyncMap) GetUint64(k string) (uint64, bool) {
	v, ok := sm.Value(k)
	if !ok {
		return 0, ok
	}
	if u, ok := zutils.InterfaceToUint(v); ok {
		return u, ok
	}
	if i, ok := zutils.InterfaceToInt(v); ok {
		return uint64(i), ok
	}
	return 0, false
}

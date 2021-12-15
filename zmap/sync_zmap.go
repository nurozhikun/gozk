package zmap

import (
	"sync"
)

// type (
// 	Key = zutils.Key
// 	Val = zutils.Val
// )

type SyncMap struct {
	*sync.Map
}

func NewSyncMap() *SyncMap {
	return &SyncMap{Map: &sync.Map{}}
}

func (sm *SyncMap) Get(key Key) (Val, bool) {
	return sm.Load(key)
}

func (sm *SyncMap) Insert(k Key, val Val) {
	sm.Store(k, val)
}

func (sm *SyncMap) Delete(key Key) (Val, bool) {
	val, ok := sm.Load(key)
	sm.Delete(key)
	return val, ok
}

func (sm *SyncMap) Value(k Key) (Val, bool) {
	return nil, false
}

func (sm *SyncMap) Len() int {
	l := 0
	sm.Range(func(k Key, v Val) bool {
		l += 1
		return true
	})
	return l
}

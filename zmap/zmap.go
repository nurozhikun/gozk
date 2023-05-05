package zmap

import (
	"sync/atomic"

	"github.com/nurozhikun/gozk/zsync"
	"github.com/nurozhikun/gozk/zutils"
)

// type Map struct {
// 	m map[Key]*Pair
// }

type (
	Key  = zutils.Key
	Val  = zutils.Val
	Pair = zutils.Pair
)

type Map map[Key]Val

func NewMap() Map {
	// return &Map{m: make(map[Key]*Pair)}
	return make(Map)
}

func (m Map) Insert(k Key, v Val) {
	m[k] = v
}

func (m Map) InsertMap(in Map) {
	for k, v := range in {
		m[k] = v
	}
}

func (m Map) InsertMapByKeys(in Map, keys ...Key) {
	if nil == in {
		return
	}
	for _, k := range keys {
		if v, ok := in[k]; ok {
			m.Insert(k, v)
		}
	}
}

func (m Map) Delete(k Key) (Val, bool) {
	v, ok := m[k]
	if ok {
		delete(m, k)
		return v, true
	} else {
		return nil, false
	}
}

func (m Map) Get(k Key) (Val, bool) {
	v, ok := m[k]
	return v, ok
}

func (m Map) Value(k Key, fnCreate func(k Key) Val) Val {
	v, ok := m[k]
	if ok {
		return v
	}
	v = fnCreate(k)
	if nil != v {
		m.Insert(k, v)
	}
	return v
}

func (m Map) Len() int {
	return len(m)
}

func (m Map) MapCall(fn func(Key, Val) (bquit bool)) {
	for k, v := range m {
		if fn(k, v) {
			break
		}
	}
}

func (m Map) ValueCall(k Key, fn func(Key, Val)) bool {
	v, b := m.Get(k)
	if b {
		fn(k, v)
	}
	return b
}

func (m Map) GetInt64(k Key) (int64, bool) {
	v, b := m.Get(k)
	if !b {
		return 0, false
	}
	if i, ok := zutils.InterfaceToInt(v); ok {
		return i, ok
	}
	if u, ok := zutils.InterfaceToUint(v); ok {
		return int64(u), ok
	}
	return 0, false
}

func (m Map) GetInt(k Key) (int, bool) {
	v, b := m.GetInt64(k)
	if b {
		return int(v), b
	}
	return 0, false
}

func (m Map) GetUint64(k Key) (uint64, bool) {
	v, ok := m[k]
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

//atomic operate v
func (m Map) TryInt64(k string, v *int64) bool {
	i, ok := m.GetInt64(k)
	if ok {
		atomic.StoreInt64(v, i)
	}
	return ok
}

func (m Map) TryAtomicInt64(k string, v *zsync.Int64) bool {
	i, ok := m.GetInt64(k)
	if ok {
		v.Set(i)
	}
	return ok
}

func (m Map) GetString(k Key) (string, bool) {
	v, ok := m[k]
	if !ok {
		return "", ok
	}
	return zutils.InterfaceToString(v)
}

func (m Map) TryAtomicString(k Key, as *zsync.String) bool {
	s, ok := m.GetString(k)
	if !ok {
		return false
	}
	as.Set(s)
	return true
}

// func (m Map) MapCall(fn func(key int, value Val) (bquit bool)) {
// 	m.Map.MapCall(func(k Key, v Val) (bquit bool) {
// 		if i, ok := k.(int); ok {
// 			return fn(i, v)
// 		}
// 		return false
// 	})
// }

// func (m Map) ValueCall(k int, fn func(int, Val)) bool {
// 	return m.Map.ValueCall(k, func(k Key, v Val) {
// 		if i, ok := k.(int); ok {
// 			fn(i, v)
// 		}
// 	})
// }

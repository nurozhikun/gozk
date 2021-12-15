package zmap

type Map struct {
	m map[Key]*Pair
}

func NewMap() *Map {
	return &Map{m: make(map[Key]*Pair)}
}

func (m *Map) Insert(k Key, v Val) {
	m.m[k] = &Pair{Key: k, Val: v}
}

func (m *Map) Delete(k Key) (Val, bool) {
	v, ok := m.m[k]
	if ok {
		delete(m.m, k)
		return v.Val, true
	} else {
		return nil, false
	}
}

func (m *Map) Get(k Key) (Val, bool) {
	v, ok := m.m[k]
	if ok {
		return v.Val, true
	} else {
		return nil, false
	}
}

func (m *Map) Value(k Key, fnCreate func(k Key) Val) Val {
	vPair, ok := m.m[k]
	if ok {
		return vPair.Val
	}
	v := fnCreate(k)
	if nil != v {
		m.Insert(k, v)
	}
	return v
}

func (m *Map) Len() int {
	return len(m.m)
}

func (m *Map) MapCall(fn func(Key, Val) (bquit bool)) {
	for k, p := range m.m {
		if fn(k, p.Val) {
			break
		}
	}
}

func (m *Map) ValueCall(k Key, fn func(Key, Val)) bool {
	v, b := m.Get(k)
	if b {
		fn(k, v)
	}
	return b
}

type IntMap struct {
	*Map
}

func NewIntMap() *IntMap {
	return &IntMap{Map: NewMap()}
}

func (m *IntMap) GetInt(key int) (int, bool) {
	v, b := m.Get(key)
	if !b {
		return 0, false
	}
	i, b := v.(int)
	return i, b
}

func (m *IntMap) MapCall(fn func(key int, value Val) (bquit bool)) {
	m.Map.MapCall(func(k Key, v Val) (bquit bool) {
		if i, ok := k.(int); ok {
			return fn(i, v)
		}
		return false
	})
}

func (m *IntMap) ValueCall(k int, fn func(int, Val)) bool {
	return m.Map.ValueCall(k, func(k Key, v Val) {
		if i, ok := k.(int); ok {
			fn(i, v)
		}
	})
}

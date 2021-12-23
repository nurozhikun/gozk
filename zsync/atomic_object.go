package zsync

import "sync/atomic"

type Int64 int64

func (i *Int64) Get() int64 {
	return atomic.LoadInt64((*int64)(i))
}

func (i *Int64) Set(v int64) {
	atomic.StoreInt64((*int64)(i), v)
}

type String struct {
	v atomic.Value
}

func (s *String) Get() string {
	if r, ok := s.v.Load().(string); ok {
		return r
	}
	return ""
}

func (s *String) Set(v string) {
	s.v.Store(v)
}

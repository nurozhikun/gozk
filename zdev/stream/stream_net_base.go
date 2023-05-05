package stream

import (
	"github.com/nurozhikun/gozk/zdev/base"
	"github.com/nurozhikun/gozk/zsync"
)

const (
	DefMaxPackSize = 1024 * 32
)

type strmNetBase struct {
	StreamBase
	ReadTimeout  zsync.Int64
	WriteTimeout zsync.Int64
	PackSize     zsync.Int64
}

func (s *strmNetBase) ISetParams(cmd *base.Command) {
	s.StreamBase.ISetParams(cmd)
	cmd.BodyMap.TryAtomicInt64(base.FieldReadTimeout, &s.ReadTimeout)
	cmd.BodyMap.TryAtomicInt64(base.FieldWriteTimeout, &s.WriteTimeout)
	cmd.BodyMap.TryAtomicInt64(base.FieldPackSize, &s.PackSize)
}

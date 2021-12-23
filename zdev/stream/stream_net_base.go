package stream

import (
	"gitee.com/sienectagv/gozk/zdev/base"
	"gitee.com/sienectagv/gozk/zsync"
)

type StrmNetBase struct {
	Stream
	readTimeout  zsync.Int64
	writeTimeout zsync.Int64
}

func (s *StrmNetBase) ISetParams(cmd *base.Command) {
	s.Stream.ISetParams(cmd)
	cmd.BodyMap.TryAtomicInt64(base.FieldReadTimeout, &s.readTimeout)
	cmd.BodyMap.TryAtomicInt64(base.FieldWriteTimeout, &s.writeTimeout)
}

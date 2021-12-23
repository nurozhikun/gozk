package vessel

import (
	"gitee.com/sienectagv/gozk/zdev/base"
	"gitee.com/sienectagv/gozk/zdev/custom"
	"gitee.com/sienectagv/gozk/zdev/stream"
)

type VesselCallback struct{}

func (cb *VesselCallback) ICreateCustom(customCode int64, cmd *Command) base.ICustom {
	return &custom.CustomBase{}
}

func (cb *VesselCallback) ICreateStream(streamCode int64, cmd *Command) base.IStream {
	switch streamCode {
	case base.Stream_TcpListener:
		return stream.CreateTcpListener(cmd.BodyMap)
	default:
		return nil
	}
}

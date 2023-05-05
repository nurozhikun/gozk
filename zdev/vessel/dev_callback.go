package vessel

import (
	"github.com/nurozhikun/gozk/zdev/base"
	"github.com/nurozhikun/gozk/zdev/custom"
	"github.com/nurozhikun/gozk/zdev/stream"
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

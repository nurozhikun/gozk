package zdev

import (
	"github.com/kataras/iris/v12"
	"github.com/nurozhikun/gozk/zdev/stream"
)

func NewStreamTcpListener() IStream {
	return &stream.StrmTcpListener{}
}

func NewStreamIrisApp(app *iris.Application) IStream {
	return stream.NewStrmIrisApp(app)
}

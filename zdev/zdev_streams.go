package zdev

import (
	"gitee.com/sienectagv/gozk/zdev/stream"
	"github.com/kataras/iris/v12"
)

func NewStreamTcpListener() IStream {
	return &stream.StrmTcpListener{}
}

func NewStreamIrisApp(app *iris.Application) IStream {
	return stream.NewStrmIrisApp(app)
}

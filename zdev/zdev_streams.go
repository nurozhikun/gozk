package zdev

import "gitee.com/sienectagv/gozk/zdev/stream"

func NewStreamTcpListener() IStream {
	return &stream.StrmTcpListener{}
}

package zdev

import (
	"github.com/nurozhikun/gozk/zdev/custom"
)

func NewCustomTcpListener(fnCreateConnCustom func() ICustom) ICustom {
	return &custom.CustomTcpListener{
		CreateConnCustom: fnCreateConnCustom,
	}
}

package zdev

import (
	"gitee.com/sienectagv/gozk/zdev/custom"
)

func NewCustomTcpListener(fnCreateConnCustom func() ICustom) ICustom {
	return &custom.CustomTcpListener{
		CreateConnCustom: fnCreateConnCustom,
	}
}

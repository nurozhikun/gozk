package custom

import (
	"net"

	"gitee.com/sienectagv/gozk/zdev/base"
	"gitee.com/sienectagv/gozk/zdev/stream"
)

type CustomTcpListener struct {
	CustomBase
	CreateConnCustom func() base.ICustom
}

func (c *CustomTcpListener) IUnpackToCommand(bin interface{}) (cmd *base.Command, err error) {
	conn, ok := bin.(net.Conn)
	if !ok {
		return
	}
	if nil == c.CreateConnCustom {
		conn.Close()
		return
	}
	cmd, err = c.CustomBase.IUnpackToCommand(bin)
	cmd.Cmd = base.Command_CreateDevice
	cmd.BodyMap.Insert(base.FieldID, c.ID()+".conn."+conn.RemoteAddr().String())
	cmd.Make().
		SetFieldStream(&stream.StrmNetConn{Conn: conn}).
		SetFieldCustom(c.CreateConnCustom())
	return
}

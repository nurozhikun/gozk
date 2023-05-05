package custom

import (
	"net"

	"github.com/nurozhikun/gozk/zdev/base"
	"github.com/nurozhikun/gozk/zdev/stream"
)

type CustomTcpListener struct {
	CustomBase
	CreateConnCustom func() base.ICustom
}

func (c *CustomTcpListener) IUnpackToCommand(bin interface{}) (cmd *base.Command, unfinished bool, err error) {
	conn, ok := bin.(net.Conn)
	if !ok {
		return
	}
	if nil == c.CreateConnCustom {
		conn.Close()
		return
	}
	cmd, unfinished, err = c.CustomBase.IUnpackToCommand(bin)
	cmd.Cmd = base.Command_CreateDevice
	cmd.BodyMap.Insert(base.FieldID, c.ID()+".conn."+conn.RemoteAddr().String())
	cmd.Make().
		SetFieldStream(&stream.StrmNetConn{Conn: conn}).
		SetFieldCustom(c.CreateConnCustom())
	return
}

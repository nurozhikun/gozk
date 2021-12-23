package custom

import (
	"net"

	"gitee.com/sienectagv/gozk/zdev/base"
)

type CustomNetTcpListener struct {
	Custom
}

func (c *CustomNetTcpListener) IUnpackToCommand(bin interface{}) (cmd *base.Command, err error) {
	conn, ok := bin.(net.Conn)
	if !ok {
		return
	}
	cmd.Cmd = base.Command_CreateDevice
	cmd, err = c.Custom.IUnpackToCommand(bin)
	cmd.BodyMap.Insert(base.FieldID, c.ID()+".conn."+conn.RemoteAddr().String())
	return
}

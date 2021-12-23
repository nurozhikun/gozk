package custom

import (
	"gitee.com/sienectagv/gozk/zdev/base"
)

type Custom struct {
	base.DevicePartner
	AckID string
}

func (c *Custom) IScheduled() (bin interface{}, unfinished bool, err error) {
	return
}

func (c *Custom) IPackCommand(cmd *base.Command) (bin interface{}, unfinished bool, err error) {
	return cmd, false, nil
}

func (c *Custom) IUnpackToCommand(bin interface{}) (cmd *base.Command, err error) {
	cmd = &base.Command{
		Cmd:        base.Command_SendData,
		ToID:       base.DeviceID_Vessel,
		BodyStruct: bin,
	}
	cmd.Make()
	cmd.ToID, _ = c.ParamsMap().GetString(base.FieldAckID)
	cmd.BodyMap.Insert(base.FieldFromID, c.ID())
	return
}

package custom

import (
	"gitee.com/sienectagv/gozk/zdev/base"
)

type CustomBase struct {
	base.DevicePartner
	// AckID string
	// base.ICallback
}

func (c *CustomBase) ISetParams(cmd *base.Command) {
	c.DevicePartner.ISetParams(cmd)
	// sl.ParamsMap().InsertMapByKeys(cmd.BodyMap, base.FieldAddr)
}

func (c *CustomBase) IScheduled() (bin interface{}, unfinished bool, err error) {
	return
}

func (c *CustomBase) IPackCommand(cmd *base.Command) (bin interface{}, unfinished bool, err error) {
	return cmd, false, nil
}

func (c *CustomBase) IUnpackToCommand(bin interface{}) (cmd *base.Command, unfinished bool, err error) {
	cmd = &base.Command{
		Cmd:  base.Command_SendData,
		ToID: base.DeviceID_Vessel,
		// BodyStruct: bin,//maybe need copy
	}
	cmd.Make()
	cmd.ToID, _ = c.ParamsMap().GetString(base.FieldAckID)
	cmd.BodyMap.Insert(base.FieldFromID, c.ID())
	return
}

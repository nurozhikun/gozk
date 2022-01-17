///
///  tag is zdev
///  fields
///  "cmd" is string
///  "id" is int
///  "code" is int
///
package zdev

import (
	"gitee.com/sienectagv/gozk/zdev/base"
	"gitee.com/sienectagv/gozk/zdev/custom"
	"gitee.com/sienectagv/gozk/zdev/stream"
	"gitee.com/sienectagv/gozk/zdev/vessel"
)

const (
	StructTag = "zdev"
	fieldCmd  = "cmd"
	fieldCode = "code"
	fieldID   = "id"
)

//value of command
const (
	Command_Unknown      = base.Command_Unknown
	Command_CreateDevice = base.Command_CreateDevice
	Command_DeleteDevice = base.Command_DeleteDevice
	Command_DevSetParams = base.Command_SetParams //
)

//value of device

type (
	Map            = base.Map
	Command        = base.Command
	Vessel         = vessel.Vessel
	FnCreateStream = base.FnCreateStream
	FnCreateCustom = base.FnCreateCustom
	ICustom        = base.ICustom
	IStream        = base.IStream
)

const (
	SKeyCmd    = base.FieldCmd    //int
	SKeyName   = base.FieldName   //string
	SKeyId     = base.FieldID     //int
	SKeyAdd    = base.FieldAddr   //string
	SKeyStream = base.FieldStream //int
	SKeyCustom = base.FieldCustom //int

	DeviceID_Vessel = base.DeviceID_Vessel
	DeviceID_Log    = base.DeviceID_Log
)

//streams
type (
	StrmEmptyWrite = stream.StrmEmptyWrite
	// StrmTcpListener = stream.StrmTcpListener
)

//customs
type (
	CustomBase        = custom.CustomBase
	CustomTcpListener = custom.CustomTcpListener
)

type VirtualNode struct {
	ID     string  `zdev:"id"`
	Addr   string  `zdev:"addr"`
	Stream IStream `zdev:"stream"`
	Custom ICustom `zdev:"custom"`
}

func CreateDevInVessel(node *VirtualNode, v *Vessel) {
	cmd := &Command{Cmd: Command_CreateDevice,
		ToID:       base.DeviceID_Vessel,
		BodyStruct: node}
	v.Dispatch(cmd)
}

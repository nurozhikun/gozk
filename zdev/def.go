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
	// Val             = zutils.Val //inteface{}
	Map             = base.Map
	Command         = base.Command
	IVessel         = base.IVessel
	IVesselCallback = base.IVesselCallback
	IDeviceObject   = base.IDeviceObject
	ICustom         = base.ICustom
	IStream         = base.IStream
)

const (
	SKeyCmd    = base.FieldCmd    //int
	SKeyName   = base.FieldName   //string
	SKeyId     = base.FieldID     //int
	SKeyAdd    = base.FieldAddr   //string
	SKeyStream = base.FieldCustom //int
	SKeyCustom = base.FieldCustom //int
)

type VirtualNode struct {
	ID     string  `zdev:"id"`
	Addr   string  `zdev:"addr"`
	Stream IStream `zdev:"stream"`
	Custom ICustom `zdev:"custom"`
}

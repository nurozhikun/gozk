package base

import (
	"gitee.com/sienectagv/gozk/zmap"
	"gitee.com/sienectagv/gozk/zreflect"
)

type (
	// Val = zutils.Val
	Map = zmap.Map
	Int = int64
)

var (
	NewMap = zmap.NewMap
)

const (
	StructTag = "zdev"
)

const (
	FieldCmd          = "cmd"    //int64
	FieldName         = "name"   //string
	FieldID           = "id"     //string
	FieldAckID        = "ack_id" //int64
	FieldFromID       = "from_id"
	FieldStream       = "stream"
	FieldCustom       = "custom"
	FieldAddr         = "addr"          //string
	FieldReadTimeout  = "read_timeout"  //int64
	FieldWriteTimeout = "write_timeout" //int64
	FieldWorkMode     = "work_mode"     //int64
	FieldChanCommand  = "chan_command"  //chan Command
)

//value of command
const (
	Command_Unknown      Int = iota
	Command_CreateDevice     //creaet a devie
	Command_DeleteDevice     //close and remove the device
	Command_RestartDevice
	// Command_RemoveDevice //only remove the device
	Command_SetParams
	Command_GetParams
	Command_SendData
)

//value of device_ID

const (
	DeviceID_Vessel = "__vessel"
	DeviceID_Log    = "__log"
)

//value of stream
const (
	Stream_Unknown     Int = 0
	Stream_TcpListener Int = 100
	Stream_TcpConnect  Int = 101 //only create by
	Stream_TcpClient   Int = 102
)

//value of WorkMode
const (
	WorkMode_Delete Int = -1
	WorkMode_Pause  Int = 0
	WorkMode_Work   Int = 1
)

type Command struct {
	Cmd        Int
	ToID       string
	BodyMap    Map
	BodyStruct interface{}
}

func (c *Command) Make() {
	if c.BodyMap == nil {
		c.BodyMap = NewMap()
	}
}

func (c *Command) StructToMap() {
	c.Make()
	c.BodyMap.InsertMap(zreflect.StructFieldsByTag(c.BodyStruct, StructTag))
}

type IVessel interface {
	Start()
	Close()
	Dispatch(cmd *Command)
	DispatchStruct(cmdCode Int, toId string, body interface{})
}

type IVesselCallback interface {
	ICreateCustom(customCode Int, cmd *Command) ICustom
	ICreateStream(streamCode Int, cmd *Command) IStream
}

type IDeviceObject interface {
	ID() string
	Vessel() IVessel
	ParamsMap() *zmap.SyncMap
	Dispatch(cmd *Command)
}

type ICustom interface {
	IDevicePartner
	//called in write goruntine
	IScheduled() (bin interface{}, unfinished bool, err error)
	IPackCommand(cmd *Command) (bin interface{}, unfinished bool, err error)
	//called in read gorunting
	IUnpackToCommand(bin interface{}) (cmd *Command, err error)
}

type IStream interface {
	IDevicePartner
	IoOpen() error
	IoClose() error
	IoCanWrite() bool
	IoWrite(bin interface{}) (err error)
	IoCanRead() bool
	IoRead() (bin interface{}, err error)
}

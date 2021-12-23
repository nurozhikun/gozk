/*zk virtual device*/
package vessel

import (
	"time"

	"gitee.com/sienectagv/gozk/zdev/base"
	"gitee.com/sienectagv/gozk/zlogger"
	"gitee.com/sienectagv/gozk/zmap"
	"gitee.com/sienectagv/gozk/zutils"
)

const (
	loopKeyVessel = "vessel"
)

type (
	Command = base.Command
	// ICallback = base.ICallback
)

type Vessel struct {
	base.FnCreateCustom
	base.FnCreateStream
	devices   base.Map
	waitGroup *zutils.LoopGroup
	chCommand chan Command
	isWorking bool
}

func (v *Vessel) Start() {
	v.devices = base.NewMap()
	v.waitGroup = zutils.NewLoopGroup()
	v.chCommand = make(chan Command, 100)
	v.isWorking = true
	v.waitGroup.GoLoop(
		loopKeyVessel,
		func() int {
			return v.chanDataProcess()
		},
		0,
		func() {},
	)
}

func (v *Vessel) Close() {
	v.waitGroup.Wait()
}

func (v *Vessel) Dispatch(cmd *Command) {
	//don't block this functions
	go func(c Command) {
		defer func() {
			if r := recover(); r != nil {
				zlogger.Info("recover in dispatch...:", r)
			}
		}()
		v.chCommand <- c
	}(*cmd)

}

//@body don't set this struct after call Dispatch
func (v *Vessel) DispatchStruct(cmdCode base.Int, toId string, body interface{}) {
	cmd := &Command{
		Cmd:        cmdCode,
		ToID:       toId,
		BodyStruct: body,
	}
	v.Dispatch(cmd)
}

func (v *Vessel) chanDataProcess() int {
DATA_LOOP:
	for {
		select {
		case cmd := <-v.chCommand:
			v.commandProcess(&cmd)
		case <-time.After(50 * time.Millisecond):
			break DATA_LOOP
		}
	}
	return 0
}

func (v *Vessel) commandProcess(cmd *Command) {
	cmd.Make().StructToMap()
	if base.DeviceID_Vessel != cmd.ToID {
		v.cmdToDevice(cmd)
	}
	switch cmd.Cmd {
	case base.Command_CreateDevice:
		v.createDevice(cmd)
	case base.Command_DeleteDevice:
		v.deleteDevice(cmd)
	case base.Command_SetParams:
		v.setParams(cmd)
	}
}

func (v *Vessel) cmdToDevice(cmd *Command) {
	i, ok := v.devices.Get(cmd.ToID)
	if !ok {
		return
	}
	d, ok := i.(*virtualDevice)
	if !ok {
		return
	}
	if d.workMode.Get() == base.Command_DeleteDevice {
		v.devices.Delete(cmd.ToID)
		return
	}
	d.Dispatch(cmd)
	return
}

//called in vessel loop
func (v *Vessel) createDevice(cmd *Command) {
	id, ok := cmd.BodyMap.GetString(base.FieldID)
	if !ok {
		return
	}
	_, ok = v.devices.Get(id)
	if ok {
		return
	}
	//create the object of virtualDevice
	vd := &virtualDevice{
		vessel:  v,
		SyncMap: zmap.NewSyncMap(),
		chCmd:   make(chan Command, 20), //must be here
		// id:      zsync.Int64(id),        //can be in init
		// sendCmdList: make([]*base.Command, 10),
	}
	vd.id.Set(id)

	//create custom
	if tmpVal, ok := cmd.BodyMap.Get(base.FieldStream); ok {
		if vd.ICustom, ok = tmpVal.(base.ICustom); !ok && nil != v.FnCreateCustom {
			if code, ok := zutils.InterfaceToInt(tmpVal); ok {
				vd.ICustom = v.FnCreateCustom(code, cmd)
			}
		}
	}
	if nil == vd.ICustom {
		return
	}
	//create stream
	if tmpVal, ok := cmd.BodyMap.Get(base.FieldCustom); ok {
		if vd.ICustom, ok = tmpVal.(base.ICustom); !ok && nil != v.FnCreateStream {
			if code, ok := zutils.InterfaceToInt(tmpVal); ok {
				vd.IStream = v.FnCreateStream(code, cmd)
			}
		}
	}
	if nil == vd.IStream {
		return
	}
	//init the device
	cmd.ToID = id
	vd.Dispatch(cmd)
	go vd.routineCommand()
	go vd.routineRead()
	v.devices.Insert(id, vd)
	zlogger.Info("device", id, "has created")
}

//called in vessel loop
func (v *Vessel) deleteDevice(cmd *Command) {
	id, ok := cmd.BodyMap.GetString(base.FieldID)
	if !ok {
		return
	}
	if d, ok := v.devices.Delete(cmd.ToID); ok {
		if dev, ok := d.(*virtualDevice); ok {
			cmd.ToID = id
			dev.Dispatch(cmd)
		}
	}
}

func (v *Vessel) setParams(cmd *Command) {

}

package vessel

import (
	"time"

	"gitee.com/sienectagv/gozk/zdev/base"
	"gitee.com/sienectagv/gozk/zlogger"
	"gitee.com/sienectagv/gozk/zmap"
	"gitee.com/sienectagv/gozk/zsync"
)

type virtualDevice struct {
	vessel *Vessel
	base.ICustom
	base.IStream
	*zmap.SyncMap
	chCmd       chan Command
	id          zsync.String
	workMode    zsync.Int64
	sendCmdList []*Command
}

//for IDeviceObject interface
// func (d *virtualDevice) Vessel() base.IVessel {
// 	return d.vessel
// }

func (d *virtualDevice) ID() string {
	return d.id.Get()
}

func (d *virtualDevice) WorkMode() int64 {
	return d.workMode.Get()
}

func (d *virtualDevice) ParamsMap() *zmap.SyncMap {
	return d.SyncMap
}

//异步发送不要阻塞
func (d *virtualDevice) Dispatch(cmd *Command) {
	if nil == cmd {
		return
	}
	if cmd.ToID == d.ID() {
		d.dispatch(cmd)
	} else {
		d.vessel.Dispatch(cmd)
	}
}

func (d *virtualDevice) Delete() {
	cmd := &Command{
		Cmd:  base.Command_DeleteDevice,
		ToID: base.DeviceID_Vessel,
	}
	cmd.Make().SetField(base.FieldID, d.ID())
	d.Dispatch(cmd)
}

//inner funcs of virtualDevice

func (d *virtualDevice) dispatch(cmd *Command) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				zlogger.Info("recover in dispatch...:", r)
			}
		}()
		d.chCmd <- *cmd
	}()
}

func (d *virtualDevice) init(cmd *Command) {
	d.ICustom.IInit(d, cmd)
	d.IStream.IInit(d, cmd)
	d.sendCmdList = make([]*Command, 10)
	cmd.BodyMap.TryAtomicString(base.FieldID, &d.id)
	d.setParams(cmd)
}

func (d *virtualDevice) open() error {
	if err := d.IoOpen(); nil != err {
		return err
	}
	return nil
}

func (d *virtualDevice) close() {
	d.sendCmdList = make([]*Command, 10)
	d.IoClose()
}

func (d *virtualDevice) setParams(cmd *Command) {
	cmd.BodyMap.TryAtomicInt64(base.FieldWorkMode, &d.workMode)
	d.ICustom.ISetParams(cmd)
	d.IStream.ISetParams(cmd)
}

func (d *virtualDevice) restart() {
	c := &Command{Cmd: base.Command_RestartDevice,
		ToID: d.ID()}
	d.Dispatch(c)
}

func (d *virtualDevice) routineCommand() {
	for d.workMode.Get() >= base.WorkMode_Pause {
		select {
		case cmd := <-d.chCmd:
			d.processCommand(&cmd)
		case <-time.After(25 * time.Millisecond):
		}
		if d.workMode.Get() <= base.WorkMode_Pause {
			d.close()
			continue
		}
		if err := d.processSending(); nil != err {
			d.restart()
		}
	}
	d.close()
	//write to io
}

func (d *virtualDevice) processCommand(cmd *Command) {
	switch cmd.Cmd {
	case base.Command_CreateDevice:
		d.init(cmd)
	case base.Command_DeleteDevice:
		d.workMode.Set(base.WorkMode_Delete)
		d.close()
	case base.Command_RestartDevice:
		d.setParams(cmd)
		d.close()
	case base.Command_SetParams:
		d.setParams(cmd)
	case base.Command_GetParams:
	case base.Command_SendData:
		d.sendCmdList = append(d.sendCmdList, cmd)
	}
}

func (d *virtualDevice) processSending() error {
	if err := d.open(); nil != err || !d.IoCanWrite() {
		return nil
	}
	//send scheduled
	err, _ := d.sendData(func() (interface{}, bool, error) {
		return d.IScheduled()
	})
	if nil != err {
		return err
	}
	if len(d.sendCmdList) == 0 {
		return nil
	}
	//write data to io
	for i := 0; i < len(d.sendCmdList); i++ {
		if !d.IoCanWrite() {
			d.sendCmdList = d.sendCmdList[i:]
			return nil
		}
		err, _ := d.sendData(func() (interface{}, bool, error) {
			return d.IPackCommand(d.sendCmdList[i])
		})
		if nil != err {
			d.sendCmdList = d.sendCmdList[i:]
			return err
		}
	}
	d.sendCmdList = make([]*Command, 10)
	return nil
}

func (d *virtualDevice) sendData(fnPack func() (interface{}, bool, error)) (err error, count int) {
	count = 0
	uf := true
	var bin interface{}
	for uf {
		if bin, uf, err = fnPack(); nil == bin || nil != err {
			break
		}
		count++
		if err = d.IoWrite(bin); nil != err {
			break
		}
	}
	if !uf {
		count = 0
	}
	return
}

func (d *virtualDevice) routineRead() {
	for d.workMode.Get() >= base.WorkMode_Pause {
		if !d.IoCanRead() {
			time.Sleep(25 * time.Millisecond)
			continue
		}
		bin, err := d.IoRead()
		if nil != err {
			d.restart()
			continue
		}
		cmd, err := d.IUnpackToCommand(bin)
		if nil != err {
			d.restart()
			continue
		}
		d.Dispatch(cmd)
	}
}

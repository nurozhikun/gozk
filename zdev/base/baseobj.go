package base

type IDevicePartner interface {
	IInit(obj IDeviceObject, cmd *Command)
	ISetParams(cmd *Command)
}

type DevicePartner struct {
	IDeviceObject
}

func (son *DevicePartner) IInit(obj IDeviceObject, cmd *Command) {
	son.IDeviceObject = obj
}

func (son *DevicePartner) ISetParams(cmd *Command) {
	son.ParamsMap().InsertMapByKeys(cmd.BodyMap,
		FieldAckID,
		FieldName,
		FieldAddr,
		FieldReadTimeout,
		FieldWriteTimeout)
}

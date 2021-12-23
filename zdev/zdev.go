package zdev

import "gitee.com/sienectagv/gozk/zdev/vessel"

func CreateVessel(cb IVesselCallback) IVessel {
	return &vessel.Vessel{Callback: cb}
}

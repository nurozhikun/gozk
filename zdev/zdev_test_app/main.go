package main

import (
	"gitee.com/sienectagv/gozk/zdev"
	"gitee.com/sienectagv/gozk/zdev/stream"
	"gitee.com/sienectagv/gozk/zdev/vessel"
	"gitee.com/sienectagv/gozk/zutils"
)

func main() {
	waitGroup := zutils.NewLoopGroup()
	vl := zdev.CreateVessel(&vessel.VesselCallback{})
	vl.Start()
	cmd := zdev.Command{}
	cmd.Cmd = zdev.Command_CreateDevice
	cmd.ID = 1
	cfgNet := stream.StrmCfgNet{}
	cmd.BodyStruct = cfgNet
	vl.Append(cmd)
	waitGroup.WaitForEnter("quit")
	vl.Close()
}

package main

import (
	"gitee.com/sienectagv/gozk/zdev"
	"gitee.com/sienectagv/gozk/zutils"
)

func main() {
	waitGroup := zutils.NewLoopGroup()
	vl := &zdev.Vessel{}
	vl.Start()
	// cmd := zdev.Command{}
	// cmd.Cmd = zdev.Command_CreateDevice

	// cfgNet := stream.StrmCfgNet{}
	// cmd.BodyStruct = cfgNet
	// vl.Append(cmd)
	waitGroup.WaitForEnter("quit")
	vl.Close()
}

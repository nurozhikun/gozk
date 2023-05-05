/*
 * @Author: wuzhikun zhikun.wu@firstack.com
 * @Date: 2023-04-18 11:09:50
 * @LastEditors: wuzhikun zhikun.wu@firstack.com
 * @LastEditTime: 2023-04-18 13:22:57
 * @Description:
 * Copyright (c) 2023 by Firstack, All Rights Reserved.
 */
package main

import (
	"time"

	"github.com/nurozhikun/gozk/zutils"
	"github.com/gonutz/w32/v2"
)

func showConsole(cmdShow int) {
	console := w32.GetConsoleWindow()
	if 0 != console {
		_, consoleProcID := w32.GetWindowThreadProcessId(console)
		if w32.GetCurrentProcessId() == consoleProcID {
			w32.ShowWindowAsync(console, cmdShow)
		}
	}
}

func main() {
	loop := zutils.NewLoopGroup()
	cmd := 0
	loop.GoLoop("test",
		func() int {
			showConsole(cmd)
			cmd = (cmd + 1) % 2
			return 1
		},
		time.Second*5,
		nil)
	loop.WaitForEnter("quit")
}

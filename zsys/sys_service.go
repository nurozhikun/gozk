/*
 * @Author: wuzhikun zhikun.wu@firstack.com
 * @Date: 2024-06-20 16:00:56
 * @LastEditors: wuzhikun zhikun.wu@firstack.com
 * @LastEditTime: 2024-06-21 13:40:58
 * @Description:
 * Copyright (c) 2024 by Firstack, All Rights Reserved.
 */
package zsys

import (
	"os"

	"github.com/kardianos/service"
	"github.com/nurozhikun/gozk/zlogger"
)

type Program struct {
	Run         func()
	Name        string
	DisplayName string
	Description string
	Executable  string
	// s   service.Service
}

func SvrWork(p *Program) (err error) {
	sr := &svr{
		p: p,
	}
	cfg := &service.Config{
		Name:        p.Name,
		DisplayName: p.DisplayName,
		Description: p.Description,
		Executable:  p.Executable,
	}
	sr.s, err = service.New(sr, cfg)
	if err != nil {
		return
	}
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			err = sr.s.Install()
			if err != nil {
				zlogger.Println(p.Name, "安装服务失败", err)
			} else {
				zlogger.Println(p.Name, "安装服务成功")
			}
			return
		case "uninstall":
			err = sr.s.Uninstall()
			if err != nil {
				zlogger.Println(p.Name, "卸载服务失败", err)
			} else {
				zlogger.Println(p.Name, "卸载服务成功")
			}
			return
		case "start":
			err = sr.s.Start()
			if err != nil {
				zlogger.Println(p.Name, "运行服务失败", err)
			} else {
				zlogger.Println(p.Name, "运行服务成功")
			}
			return
		case "stop":
			err = sr.s.Stop()
			if err != nil {
				zlogger.Println(p.Name, "停止服务失败", err)
			} else {
				zlogger.Println(p.Name, "停止服务成功")
			}
			return
		}
	}
	err = sr.s.Run()
	return
}

// //call this
// func (p *Service) WorkInMain() {

// }

type svr struct {
	p *Program
	s service.Service
}

func (p *svr) Start(s service.Service) error {
	zlogger.Println(p.p.Name, "service is running ...")
	go p.p.Run()
	return nil
}

func (p *svr) Stop(s service.Service) error {
	zlogger.Println(p.p.Name, "service is stopped")
	return nil
}

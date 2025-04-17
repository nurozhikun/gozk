/*
 * @Author: wuzhikun zhikun.wu@firstack.com
 * @Date: 2023-04-18 10:23:10
 * @LastEditors: wuzhikun zhikun.wu@firstack.com
 * @LastEditTime: 2023-09-25 14:13:28
 * @Description:
 * Copyright (c) 2023 by Firstack, All Rights Reserved.
 */
package zcfg

import (
	"github.com/nurozhikun/gozk/zlogger"
	"github.com/nurozhikun/gozk/zsys"
	"github.com/nurozhikun/gozk/zutils"
	"gopkg.in/ini.v1"
)

type CfgRedis struct {
	AddrTcp     string `ini:"addr_tcp"`
	AddrUrl     string `ini:"addr_url"`
	MaxActive   int    `ini:"max_active"`
	IdleTimeout int    `ini:"idle_timeout"`
	SlaveTcp    string `ini:"slave_tcp"`
	SlaveUrl    string `ini:"slave_url"`
	User        string `ini:"user"`
	Password    string `ini:"password"`
}

type CfgServer struct {
	AddrTcp string `ini:"addr_tcp"`
	AddrUrl string `ini:"addr_url"`
}

type CfgDb struct{}

func IniMapToCfg(cfg interface{}) error {
	defFile := zutils.RootFileExSuffix()
	files := []interface{}{}
	file := defFile + "-user.ini" //{appname}-user.ini
	// zlogger.Println(file)
	if zsys.FileExist(file) {
		files = append(files, file)
	}
	file = defFile + ".ini" //{appname}.ini
	// zlogger.Println(file)
	sections, err := ini.Load(file, files...)
	if nil != err {
		zlogger.Error(err)
		return err
	}
	// tcfg := &Config{}
	// sections.MapTo(tcfg)
	// zlogger.Println(tcfg.Code)
	err = sections.MapTo(cfg)
	if err != nil {
		zlogger.Error(err)
		return err
	}
	return nil
}

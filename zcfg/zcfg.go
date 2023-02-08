package zcfg

import (
	"gitee.com/sienectagv/gozk/zlogger"
	"gitee.com/sienectagv/gozk/zsys"
	"gitee.com/sienectagv/gozk/zutils"
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

type CfgDb struct{}

func IniMapToCfg(cfg interface{}) error {
	defFile := zutils.RootFileExSuffix()
	files := []interface{}{}
	file := defFile + ".ini" //[appname].ini
	// zlogger.Println(file)
	if zsys.FileExist(file) {
		files = append(files, file)
	}
	file = defFile + "-def.ini" //[appname]-def.ini
	// zlogger.Println(file)
	sections, err := ini.Load(file, files...)
	if nil != err {
		zlogger.Error(err)
		return nil
	}
	// tcfg := &Config{}
	// sections.MapTo(tcfg)
	// zlogger.Println(tcfg.Code)
	sections.MapTo(cfg)
	return nil
}

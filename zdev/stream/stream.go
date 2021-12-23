package stream

import (
	"gitee.com/sienectagv/gozk/zdev/base"
)

type StrmCfgNet struct {
	Addr     string `zdev:"addr"`
	UserName string `zdev:"user_name"`
	Password string `zdev:"password"`
}

type Stream struct {
	base.DevicePartner
}

type StrmEmptyWrite struct{}

func (s *StrmEmptyWrite) IoCanWrite() bool {
	return true
}

func (s *StrmEmptyWrite) IoWrite(bin interface{}) (err error) {
	return nil
}

// func (s *Stream) Write(bin interface{}) (err error) {
// 	return nil
// }

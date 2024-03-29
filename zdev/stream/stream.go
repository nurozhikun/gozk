package stream

import (
	"github.com/nurozhikun/gozk/zdev/base"
)

type StrmCfgNet struct {
	Addr     string `zdev:"addr"`
	UserName string `zdev:"user_name"`
	Password string `zdev:"password"`
}

type StreamBase struct {
	base.DevicePartner
}

type StrmEmptyWrite struct{}

func (s *StrmEmptyWrite) IoCanWrite() bool {
	return true
}

func (s *StrmEmptyWrite) IoWrite(bin interface{}) (err error) {
	return nil
}

// func (s *StreamBase) Write(bin interface{}) (err error) {
// 	return nil
// }

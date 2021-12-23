package stream

import (
	"errors"
	"net"

	"gitee.com/sienectagv/gozk/zdev/base"
)

type StrmTcpListener struct {
	StrmNetBase
	StrmEmptyWrite
	listener *net.TCPListener
}

func CreaeTcpListener(m base.Map) *StrmTcpListener {
	return &StrmTcpListener{}
}

func (sl *StrmTcpListener) ISetParams(cmd *base.Command) {
	sl.StrmNetBase.ISetParams(cmd)
	sl.ParamsMap().InsertMapByKeys(cmd.BodyMap, base.FieldAddr)
}

func (sl *StrmTcpListener) IoOpen() (err error) {
	if nil != sl.listener {
		return
	}
	addr, ok := sl.ParamsMap().GetString(base.FieldAddr)
	if !ok {
		return errors.New("this no address for listener")
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if nil != err {
		return err
	}
	sl.listener, err = net.ListenTCP("tcp", tcpAddr)
	return
}

func (sl *StrmTcpListener) IoClose() error {
	if nil != sl.listener {
		sl.listener.Close()
		sl.listener = nil
	}
	return nil
}

func (sl *StrmTcpListener) IoCanRead() bool {
	return nil != sl.listener
}

func (sl *StrmTcpListener) IoRead() (bin interface{}, err error) {
	if nil != sl.listener {
		return
	}
	conn, err := sl.listener.AcceptTCP()
	return conn, err
}

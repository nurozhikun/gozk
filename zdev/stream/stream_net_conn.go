package stream

import (
	"errors"
	"net"
	"time"
)

type StrmNetConn struct {
	strmNetBase
	net.Conn
	left []byte
}

func (s *StrmNetConn) IoOpen() error {
	if nil == s.Conn {
		return errors.New("the conn is nil")
	}
	return nil
}

func (s *StrmNetConn) IoClose() error {
	if nil != s.Conn {
		s.Conn.Close()
	}
	return nil
}

func (s *StrmNetConn) IoCanWrite() bool {
	if nil == s.Conn {
		return false
	}
	b, err := s.write(s.left)
	return (b && nil == err)
}

func (s *StrmNetConn) IoWrite(bin interface{}) error {
	if nil == bin {
		return nil
	}
	buff, ok := bin.([]byte)
	if !ok {
		return nil
	}
	if 0 != s.WriteTimeout.Get() {
		s.SetDeadline(time.Now().Add(time.Duration(s.WriteTimeout.Get()) * time.Millisecond))
	}
	_, err := s.write(buff)
	return err
}

func (s *StrmNetConn) write(buff []byte) (finished bool, err error) {
	n, err := s.Write(buff)
	if n != len(buff) {
		s.left = buff[n:]
	} else {
		finished = true
	}
	if nil != err { //delete the device from vessel
		if e, ok := err.(net.Error); ok && e.Timeout() {
			err = nil
		}
	}
	if nil != err {
		s.Delete()
	}
	return
}

func (s *StrmNetConn) IoCanRead() bool {
	return nil != s.Conn
}

func (s *StrmNetConn) IoRead() (bin interface{}, err error) {
	size := s.PackSize.Get()
	if size == 0 {
		size = DefMaxPackSize
	}
	buff := make([]byte, size)
	n, err := s.Read(buff)
	return buff[0:n], err
}

package ztest

import (
	"fmt"
	"net"
	"testing"
)

func TestConn(t *testing.T) {
	var i interface{} = &net.TCPConn{}
	c, ok := i.(net.Conn)
	fmt.Println(c, ok)
}

func TestType(t *testing.T) {
	s := "s12"
	var i interface{} = s
	n, _ := i.(int)
	fmt.Println(n)
}

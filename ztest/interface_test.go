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

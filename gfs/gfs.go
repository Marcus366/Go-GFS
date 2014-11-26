package gfs

import (
	"fmt"
	"net"
	"net/rpc"
)

var (
	MasterIP = net.ParseIP("127.0.0.1")
	Conn     *rpc.Client
)

func init() {
	addr := fmt.Sprintf("%s:%v", MasterIP.String(), OpenClosePort)
	Conn, _ = rpc.Dial("tcp", addr)
}

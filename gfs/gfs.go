package gfs

import (
  "fmt"
  "net"
  "net/rpc"

  "github.com/sysu2012zzp/Go-GFS/transport"
)

var (
  MasterIP = net.ParseIP("127.0.0.1")
  Conn *rpc.Client
)

func init() {
  addr := fmt.Sprintf("%s:%v", MasterIP.String(), transport.OpenClosePort)
  Conn, _ = rpc.Dial("tcp", addr)
}

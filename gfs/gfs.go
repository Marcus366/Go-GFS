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

func Open(fullFileName string, flag int64) (int32, error) {
  args := transport.OpenArgs { fullFileName, flag }
  var reply transport.OpenReply
  err := Conn.Call("OpenClose.Open", &args, &reply)
  return reply.FD, err
}

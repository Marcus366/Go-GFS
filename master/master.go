package master

import (
  "container/list"
  "fmt"
  "log"
  "net"
  "net/rpc"

  "github.com/sysu2012zzp/Go-GFS/transport"
)

type Master struct {
  chunkServers *list.List
}

func NewMaster() *Master {
  return &Master {
    chunkServers: list.New() }
}

func (m *Master) Main() {
  go m.openHeartbeatServer()
  go m.openRegisterServer()
}

func (m *Master) openHeartbeatServer() {
  r := rpc.NewServer()
  r.Register(&transport.Heartbeat{m})

  l, e := net.Listen("tcp", ":2345")
  if e != nil {
    log.Fatal("listen error: ", e)
  }
  go r.Accept(l)
}

func (m *Master) openRegisterServer() {
  r := rpc.NewServer()
  r.Register(&transport.Reg{m})

  l, e := net.Listen("tcp", ":2346")
  if e != nil {
    log.Fatal("open register server fail:", e)
  }
  go r.Accept(l)
}

func (m *Master) KeepAlive(args *transport.HeartbeatArgs, reply *transport.HeartbeatReply) error {
  if args.IP != nil {
    fmt.Println("Heartbeat IP:", args.IP)
  }
  return nil
}

func (m *Master) Register(args *transport.RegArgs, reply *transport.RegReply) error {
  fmt.Println("Register:", args.IP)
  return nil
}

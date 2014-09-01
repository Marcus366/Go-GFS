package master

import (
  "container/list"
  "fmt"
  "log"
  "net"
  "net/rpc"
  "strings"

  "github.com/sysu2012zzp/Go-GFS/transport"
)

type Master struct {
  chunkServers *list.List
}

func NewMaster() *Master {
  return &Master { chunkServers: list.New() }
}

func (m *Master) Main() {
  m.openHeartbeatServer()
  m.openRegisterServer()
  m.openManagerServer()
}

func (m *Master) openHeartbeatServer() {
  r := rpc.NewServer()
  r.Register(&transport.Heartbeat{m})

  addr := fmt.Sprintf(":%v", transport.HeartbeatPort)
  l, e := net.Listen("tcp", addr)
  if e != nil {
    log.Fatal("listen error: ", e)
  }
  go r.Accept(l)
}

func (m *Master) openRegisterServer() {
  r := rpc.NewServer()
  r.Register(&transport.Reg{m})

  addr := fmt.Sprintf(":%v", transport.RegisterPort)
  l, e := net.Listen("tcp", addr)
  if e != nil {
    log.Fatal("open register server fail:", e)
  }
  go r.Accept(l)
}

func (m *Master) openManagerServer() {
  r := rpc.NewServer()
  r.Register(&transport.Manager{m})

  addr := fmt.Sprintf(":%v", transport.OpenClosePort)
  l, e := net.Listen("tcp", addr)
  if e != nil {
    log.Fatal("open openclose server fail:", e)
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
  chunk := &ChunkServer{ IP: args.IP }
  m.chunkServers.PushFront(chunk)
  return nil
}

func (m *Master) OpenFile(name string, flag int, perm transport.FileMode) (int32, error) {
  if flag & O_CREATE != 0 {
    slice := strings.Split(name, '/')
  }
  fmt.Println("OpenFile: ", name);
  return 0, nil
}

func (m *Master) Open(name string) (int32, error) {
  fmt.Println("Open: ", name)
  return 0, nil
}

func (m *Master) Close(fd int32) error {
  fmt.Println("Close: ", fd)
  return nil
}

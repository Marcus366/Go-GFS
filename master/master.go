package master

import (
  "container/list"
  "fmt"
  "net"
)

type Master struct {
  chunkServers *list.List
}

func NewMaster() *Master {
  return &Master {
    chunkServers: list.New() }
}

func (m *Master) Main() {
  listner, err := net.Listen("tcp", ":4399")
  if err != nil {
    fmt.Println("[master] listen error: ", err)
  }
  for {
    conn, err := listner.Accept()
    if err != nil {
      fmt.Println("[master] accept error: ", err)
    }
    _ = conn
    //TODO: go Handle conn
    //go Handler(conn)
  }
}

func Handler(conn net.Conn) {
}

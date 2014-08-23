package gfs

import (
  "github.com/sysu2012zzp/Go-GFS/transport"
)

type File struct {
}

func Open(name string) (int32, error) {
  args := transport.OpenArgs { name, 0 }
  var reply transport.OpenReply
  err := Conn.Call("Manager.Open", &args, &reply)
  return reply.Fd, err
}


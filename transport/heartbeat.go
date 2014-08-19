package transport

import (
  "net"
  "fmt"
)

type Heartbeat struct {
}

type HeartbeatArgs struct {
  IP net.IP
}

type HeartbeatReply struct {
  hey bool
}

func (h *Heartbeat) KeepAlive(args *HeartbeatArgs, reply *HeartbeatReply) error {
  fmt.Println("Heartbeat")
  reply.hey = true
  if args.IP != nil {
    fmt.Println("Heartbeat: ", args.IP.String())
  }
  return nil
}

package transport

import (
  "net"
)

type HeartbeatCallback interface {
  KeepAlive(args *HeartbeatArgs, reply *HeartbeatReply) error
}

type HeartbeatArgs struct {
  IP net.IP
}

type HeartbeatReply struct {
}

type Heartbeat struct {
  Callback HeartbeatCallback
}

func (h *Heartbeat) KeepAlive(args *HeartbeatArgs, reply *HeartbeatReply) error {
  if h.Callback != nil {
    return h.Callback.KeepAlive(args, reply)
  }
  return nil
}

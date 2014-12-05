package common

import (
	"net"
)

type HeartbeatArgs struct {
	IP   net.IP
	Port int
}

type HeartbeatReply struct {
}

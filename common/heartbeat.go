package common

import (
	"net"
)

type HeartbeatArgs struct {
	IP net.IP
}

type HeartbeatReply struct {
}

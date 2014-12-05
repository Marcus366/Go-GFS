package master

import (
	"net"
)

type ChunkServer struct {
	IP   net.IP
	Port int
}

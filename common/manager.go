package common

import (
	"net"
)

const (
	O_RDONLY = 0x0
	O_WRONLY = 0x1
	O_RDWR   = 0x2
	O_APPEND = 0x400
	O_CREATE = 0x40

	/* the following flag maybe supported later
	 * O_EXCL
	 * O_SYNC
	 * O_TRUNC
	 */
)

type OpenArgs struct {
	Name string
	Flag int
	Perm uint32
}

type OpenReply struct {
	Fd int32
	Err error
}

type CloseArgs struct {
	Fd int32
}

type CloseReply struct {
}

type WriteArgs struct {
	Fd  int32
	Off int64
}

type WriteTempReply struct {
	IP   net.IP
	Uuid uint64
	Size uint64
	Err error
}

type WriteTempArgs struct {
	Uuid uint64
	Buf []byte
	Off  int64
}

type WriteReply struct {
	Bytes int
	Err   error
}

type ReadArgs struct {
	Bytes int
	Off   int64
}

type ReadReply struct {
	Buf []byte
	Err error
}

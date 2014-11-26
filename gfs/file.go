package gfs

import (
	"os"

	"github.com/sysu2012zzp/Go-GFS/transport"
)

type FileMode os.FileMode

type File struct {
}

func OpenFile(name string, flag int32, perm FileMode) (int32, error) {
	args := transport.OpenArgs{name, 0, 0}
	var reply transport.OpenReply
	err := Conn.Call("Manager.OpenFile", &args, &reply)
	return reply.Fd, err
}

func Open(name string) (int32, error) {
	args := transport.OpenArgs{name, 0, 0}
	var reply transport.OpenReply
	err := Conn.Call("Manager.Open", &args, &reply)
	return reply.Fd, err
}

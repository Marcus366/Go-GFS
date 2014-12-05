package gfs

import (
	"fmt"
	"net/rpc"
	"runtime"

	"GoFS/common"
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

type FileMode uint32

const (
	ModeDir FileMode = 1 << (32 - 1 - iota)
	ModeAppend
	ModeExlusive
	ModeTemporary
	ModeSymlink
)

type File struct {
	Fd int32
}

func OpenFile(name string, flag int, perm FileMode) (*File, error) {
	args := common.OpenArgs{name, flag, uint32(perm)}
	var reply common.OpenReply
	err := Conn.Call("Master.OpenFile", &args, &reply)
	return &File{reply.Fd}, err
}

func Open(name string) (*File, error) {
	args := common.OpenArgs{name, 0, 0}
	var reply common.OpenReply
	err := Conn.Call("Master.Open", &args, &reply)
	return &File{reply.Fd}, err
}

func (f *File) Close() error {
	return nil
}

func (f *File) Read(b []byte) (int, error) {
	return 0, nil
}

func (f *File) ReadAt(b []byte, off int64) (int, error) {
	return 0, nil
}

func (f *File) Write(b []byte) (int, error) {
	runtime.GOMAXPROCS(2)
	args := common.WriteArgs{f.Fd, -1}
	var reply common.WriteTempReply
	fmt.Println("Write")
	err := Conn.Call("Master.Write", &args, &reply)
	if err != nil {
		fmt.Println("Call Master Failed:", err)
		return 0, err
	}

	fmt.Println("Call Master Return IP:", reply.IP.String(), "Port:", reply.Port)
	addr := fmt.Sprintf("%s:%v", reply.IP.String(), reply.Port)
	fmt.Println("Connect ChunkServer Addr:", addr)
	conn, err := rpc.Dial("tcp", addr)
	if err != nil {
		return 0, err
	}

	arg := &common.WriteTempArgs{reply.Uuid, b, -1}
	var reply2 common.WriteReply
	err = conn.Call("ChunkServer.Write", &arg, &reply2)
	fmt.Println("Call Write to ChunkServer")

	return reply2.Bytes, err
}

func (f *File) WriteAt(b []byte, off int64) (int, error) {
	return 0, nil
}

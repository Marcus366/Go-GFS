package gfs

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
}

func OpenFile(name string, flag int32, perm FileMode) (int32, error) {
	args := OpenArgs{name, 0, 0}
	var reply OpenReply
	err := Conn.Call("Master.OpenFile", &args, &reply)
	return reply.Fd, err
}

func Open(name string) (int32, error) {
	args := OpenArgs{name, 0, 0}
	var reply OpenReply
	err := Conn.Call("Master.Open", &args, &reply)
	return reply.Fd, err
}

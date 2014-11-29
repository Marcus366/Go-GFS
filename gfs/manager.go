package gfs

type OpenArgs struct {
	Name string
	Flag int
	Perm FileMode
}

type OpenReply struct {
	Fd int32
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
	Msg ChunkServerMsg
	Uuid int64
	Size int64
	Err error
}

type WriteTempArgs struct {
	Uuid int64
	Buf []byte
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

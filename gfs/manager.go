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

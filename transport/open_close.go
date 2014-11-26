package transport

import (
	"errors"
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

type ManagerCallback interface {
	OpenFile(name string, flag int, perm FileMode) (int32, error)
	Open(name string) (int32, error)
	Close(fd int32) error
}

type Manager struct {
	Callback ManagerCallback
}

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

func (oc *Manager) OpenFile(args *OpenArgs, reply *OpenReply) error {
	if oc.Callback != nil {
		fd, err := oc.Callback.OpenFile(args.Name, args.Flag, args.Perm)
		reply.Fd = fd
		return err
	} else {
		return errors.New("Callback nil")
	}
}

func (oc *Manager) Open(args *OpenArgs, reply *OpenReply) error {
	if oc.Callback != nil {
		fd, err := oc.Callback.Open(args.Name)
		reply.Fd = fd
		return err
	} else {
		return errors.New("Callback nil")
	}
}

func (oc *Manager) Close(args *CloseArgs, reply *CloseReply) error {
	if oc.Callback != nil {
		err := oc.Callback.Close(args.Fd)
		return err
	} else {
		return errors.New("Callback nil")
	}
}

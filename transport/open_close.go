package transport

import (
  "errors"
)

type ManagerCallback interface {
  Open(name string) (int32, error)
  Close(fd int32) error
}

type Manager struct {
  Callback ManagerCallback
}

type OpenArgs struct {
  Name string
  Flag  int64
}

type OpenReply struct {
  Fd int32
}

type CloseArgs struct {
  Fd int32
}

type CloseReply struct {
}

func (oc *Manager) Open(args *OpenArgs, reply *OpenReply) error {
  if oc.Callback != nil {
    fd, err := oc.Callback.Open(args.Name)
    reply.Fd = fd;
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

package transport

import (
  "errors"
)

type OpenCloseCallback interface {
  Open(fullFileName string, flag int64) (int32, error)
  Close(fd int32) error
}

type OpenClose struct {
  Callback OpenCloseCallback
}

type OpenArgs struct {
  FullFileName string
  Flag  int64
}

type OpenReply struct {
  FD int32
}

type CloseArgs struct {
  FD int32
}

type CloseReply struct {
}

func (oc *OpenClose) Open(args *OpenArgs, reply *OpenReply) error {
  if oc.Callback != nil {
    fd, err := oc.Callback.Open(args.FullFileName, args.Flag)
    reply.FD = fd
    return err
  } else {
    return errors.New("OpenClose Callback nil")
  }
}

func (oc *OpenClose) Close(args *CloseArgs, reply *CloseReply) error {
  if oc.Callback != nil {
    err := oc.Callback.Close(args.FD)
    return err
  } else {
    return errors.New("OpenClose Callback nil")
  }

}

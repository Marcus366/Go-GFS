package transport

import (
	"net"
)

type RegCallback interface {
	Register(args *RegArgs, reply *RegReply) error
}

type RegArgs struct {
	IP net.IP
}

type RegReply struct {
}

type Reg struct {
	Callback RegCallback
}

func (r *Reg) Register(args *RegArgs, reply *RegReply) error {
	if r.Callback != nil {
		return r.Callback.Register(args, reply)
	}
	return nil
}

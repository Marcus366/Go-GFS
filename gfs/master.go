package gfs

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	//"strings"
)

type Master struct {
	ChunkServers map[string]*ChunkServerMsg
}

func NewMaster() *Master {
	m := new(Master)
	m.ChunkServers = make(map[string]*ChunkServerMsg)

	return m
}

func (m *Master) Main() {
	m.openHeartbeatServer()
	m.openManagerServer()
}

func (m *Master) openHeartbeatServer() {
	r := rpc.NewServer()
	r.Register(m)

	addr := fmt.Sprintf(":%v", HeartbeatPort)
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatal("listen error: ", e)
	}
	go r.Accept(l)
}

func (m *Master) openManagerServer() {
	r := rpc.NewServer()
	r.Register(m)

	addr := fmt.Sprintf(":%v", OpenClosePort)
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatal("open openclose server fail:", e)
	}
	go r.Accept(l)
}

func (m *Master) KeepAlive(args *HeartbeatArgs, reply *HeartbeatReply) error {
	if args.IP != nil {
		fmt.Println("Heartbeat IP:", args.IP)
		ip := args.IP.String()
		if m.ChunkServers[ip] == nil {
			cs := new(ChunkServerMsg)
			cs.IP = args.IP
			m.ChunkServers[ip] = cs
		}
	}
	return nil
}

func (m *Master) OpenFile(args *OpenArgs, reply *OpenReply) error {
	//if flag&O_CREATE != 0 {
		//slice := strings.Split(name, "/")
	//}
	fmt.Println("OpenFile: ", args.Name)
	return nil
}

func (m *Master) Open(args *OpenArgs, reply *OpenReply) error {
	fmt.Println("Open: ", args.Name)
	return nil
}

func (m *Master) Close(args *CloseArgs, reply *CloseReply) error {
	fmt.Println("Close: ", args.Fd)
	return nil
}

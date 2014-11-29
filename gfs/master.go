package gfs

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	//"strings"
)

type Master struct {
	chunkServers map[string]*ChunkServerMsg
	openFiles    []*FileMsg
	nameSpace    *Namespace
}

func NewMaster() *Master {
	m := new(Master)
	m.chunkServers = make(map[string]*ChunkServerMsg)
	m.openFile = make([]*FileMsg, 1024)
	m.nameSpace = NewNamespace()

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
		if m.chunkServers[ip] == nil {
			cs := new(ChunkServerMsg)
			cs.IP = args.IP
			m.chunkServers[ip] = cs
		}
	}
	return nil
}

func (m *Master) OpenFile(args *OpenArgs, reply *OpenReply) error {
	if args.Flag&O_CREATE != 0 {
		_, err := m.nameSpace.createFile(args.Name, args.Flag, args.Perm)
		if err != nil {
			fmt.Println("open file: ", args.Name, " fail.")
			return err
		}
	}
	fmt.Println("OpenFile: ", args.Name)
	return nil
}

func (m *Master) Open(args *OpenArgs, reply *OpenReply) error {
	fmt.Println("Open: ", args.Name)
	filemsg, err := m.nameSpace.findFile(args.Name)
	if err != nil {
		reply.Err = err
		return nil
	}

	for i, msg := range m.openFiles {
		if msg == nil {
			m.openFiles[i] = filemsg
			reply.Fd = i
			return nil
		}
	}
	return nil
}

func (m *Master) Close(args *CloseArgs, reply *CloseReply) error {
	fmt.Println("Close: ", args.Fd)
	m.openFiles[args.Fd] = nil
	return nil
}

func (m *Master) Write(args *WriteArgs, reply *WriteTempReply) error {
	msg := m.openFiles[args.Fd]
	if msg == nil {
		reply.Err = errors.New("The file has not been opened")
		return nil
	}

	if args.Off == -1 {
		lastChunk := msg.chunks.Back().Value.(*ChunkMsg)
		reply.Msg = *lastChunk.location
		reply.Uuid = lastChunk.uuid
		reply.Size = lastChunk.size
		return nil
	}

	return nil
}

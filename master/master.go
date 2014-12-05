package master

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	//"strings"
	"GoFS/common"
)

type Master struct {
	chunkServers map[string]*ChunkServer
	openFiles    []*File
	nameSpace    *Namespace
}

func NewMaster() *Master {
	m := new(Master)
	m.chunkServers = make(map[string]*ChunkServer)
	m.openFiles = make([]*File, 1024)
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

	addr := fmt.Sprintf(":%v", common.HeartbeatPort)
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatal("listen error: ", e)
	}
	go r.Accept(l)
}

func (m *Master) openManagerServer() {
	r := rpc.NewServer()
	r.Register(m)

	addr := fmt.Sprintf(":%v", common.ManagerPort)
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatal("open manager server fail:", e)
	}
	go r.Accept(l)
}

func (m *Master) KeepAlive(args *common.HeartbeatArgs, reply *common.HeartbeatReply) error {
	if args.IP != nil {
		fmt.Println("Heartbeat IP:", args.IP, " Port:", args.Port)
		ip := args.IP.String()
		if m.chunkServers[ip] == nil {
			cs := new(ChunkServer)
			cs.IP = args.IP
			cs.Port = args.Port
			m.chunkServers[ip] = cs
		}
	}
	return nil
}

func (m *Master) OpenFile(args *common.OpenArgs, reply *common.OpenReply) error {
	if args.Flag&common.O_CREATE != 0 {
		file, err := m.nameSpace.createFile(args.Name, args.Flag, args.Perm)
		if err != nil {
			fmt.Println("open file: ", args.Name, " fail.")
			return err
		}

		for _, cs := range m.chunkServers {
				c := NewChunk(cs)
				file.chunks.PushBack(c)
				break
		}

		for i, msg := range m.openFiles {
			if msg == nil {
				m.openFiles[i] = file
				reply.Fd = int32(i)
			}
		}
		
		return nil
	}

	fmt.Println("OpenFile: ", args.Name)
	return nil
}

func (m *Master) Open(args *common.OpenArgs, reply *common.OpenReply) error {
	fmt.Println("Open: ", args.Name)
	filemsg, err := m.nameSpace.findFile(args.Name)
	if err != nil {
		return err
	}

	for i, msg := range m.openFiles {
		if msg == nil {
			m.openFiles[i] = filemsg
			reply.Fd = int32(i)
			return nil
		}
	}
	return nil
}

func (m *Master) Close(args *common.CloseArgs, reply *common.CloseReply) error {
	fmt.Println("Close: ", args.Fd)
	m.openFiles[args.Fd] = nil
	return nil
}

func (m *Master) Write(args *common.WriteArgs, reply *common.WriteTempReply) error {
	fmt.Println("Write fd:", args.Fd, "Offset:", args.Off)
	file := m.openFiles[args.Fd]
	if file == nil {
		return errors.New("The file has not been opened")
	}

	if args.Off == -1 {
		lastChunk := file.chunks.Back().Value.(*Chunk)
		reply.IP = lastChunk.location.IP
		reply.Port = lastChunk.location.Port
		reply.Uuid = lastChunk.uuid
		reply.Size = lastChunk.size
		fmt.Println("Write Return IP", reply.IP.String(), "Port", reply.Port, "Uuid", reply.Uuid, "Size", reply.Size)
		return nil
	}

	return nil
}

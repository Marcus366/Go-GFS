package gfs

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

type ChunkServer struct {
	LocalIP  net.IP
	MasterIP net.IP
}

func NewChunkServer(ip net.IP) *ChunkServer {
	cs := new(ChunkServer)

	cs.LocalIP = LocalIP()
	cs.MasterIP = ip

	return cs
}

func (cs *ChunkServer) Main(exitChan chan string) {
	go cs.sendHeartbeat(exitChan)
}

func (cs *ChunkServer) sendHeartbeat(exitChan chan string) {
	client, err := rpc.Dial("tcp", cs.MasterIP.String()+":2345")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := HeartbeatArgs{IP: LocalIP()}
	fmt.Println("IP:", args.IP.String())
	var reply HeartbeatReply
	for {
		time.Sleep(time.Second * 10)
		err := client.Call("Master.KeepAlive", &args, &reply)
		if err != nil {
			fmt.Println("rpc call failed:", err)
			exitChan <- err.Error()
			return
		}
	}
}

func (cs *ChunkServer) Write(args *WriteTempArgs, reply *WriteReply) error {
	return nil
}

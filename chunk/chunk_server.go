package chunk

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/rpc"
	"os"
	"strconv"
	"time"

	"GoFS/common"
)

type ChunkServer struct {
	LocalIP  net.IP
	MasterIP net.IP
	Port     int
}

func NewChunkServer(ip net.IP) *ChunkServer {
	cs := new(ChunkServer)

	cs.LocalIP = common.LocalIP()
	cs.MasterIP = ip

	return cs
}

func (cs *ChunkServer) Main(exitChan chan string) {
	r := rpc.NewServer()
	r.Register(cs)

loop:
	port := 1024 + rand.Intn(65536-1024)
	addr := fmt.Sprintf(":%v", port)
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatal("open manager server fail:", e)
		goto loop
	}
	cs.Port = port
	go r.Accept(l)

	go cs.sendHeartbeat(exitChan)
}

func (cs *ChunkServer) sendHeartbeat(exitChan chan string) {
	client, err := rpc.Dial("tcp", cs.MasterIP.String()+":2345")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := common.HeartbeatArgs{IP: cs.LocalIP, Port: cs.Port}
	var reply common.HeartbeatReply
	fmt.Println("IP:", args.IP.String())

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

func (cs *ChunkServer) Write(args *common.WriteTempArgs, reply *common.WriteReply) error {
	fmt.Println("Write Uuid:", args.Uuid, "Content:", string(args.Buf))
	name := strconv.Itoa(int(args.Uuid))
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	reply.Bytes, err = file.Write(args.Buf)
	return err
}

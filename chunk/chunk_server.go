package chunk

import (
	"os"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"strconv"
	"time"

	"GoFS/common"
)

type ChunkServer struct {
	LocalIP  net.IP
	MasterIP net.IP
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

	addr := fmt.Sprintf(":%v", common.ManagerPort)
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatal("open manager server fail:", e)
	}
	go r.Accept(l)

	go cs.sendHeartbeat(exitChan)
}

func (cs *ChunkServer) sendHeartbeat(exitChan chan string) {
	client, err := rpc.Dial("tcp", cs.MasterIP.String()+":2345")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := common.HeartbeatArgs{IP: cs.LocalIP}
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
	name := strconv.Itoa(int(args.Uuid))
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	reply.Bytes, reply.Err = file.Write(args.Buf)
	return nil
}

package chunk

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"

	"github.com/sysu2012zzp/Go-GFS/transport"
	"github.com/sysu2012zzp/Go-GFS/utils"
)

type Chunk struct {
	LocalIP  net.IP
	MasterIP net.IP
}

func NewChunk(ip net.IP) *Chunk {
	return &Chunk{
		LocalIP:  utils.LocalIP(),
		MasterIP: ip
	}
}

func (c *Chunk) Main(exitChan chan string) {
	err := c.register()
	if err != nil {
		exitChan <- err.Error()
		return
	}
	go c.sendHeartbeat(exitChan)
}

func (c *Chunk) register() error {
	client, err := rpc.Dial("tcp", c.MasterIP.String()+":2346")
	if err != nil {
		log.Fatal("register fail:", err)
	}
	args := transport.RegArgs{IP: c.LocalIP}
	var reply transport.RegReply
	return client.Call("Reg.Register", &args, &reply)
}

func (c *Chunk) sendHeartbeat(exitChan chan string) {
	client, err := rpc.Dial("tcp", c.MasterIP.String()+":2345")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := transport.HeartbeatArgs{IP: utils.LocalIP()}
	fmt.Println("IP:", args.IP.String())
	var reply transport.HeartbeatReply
	for {
		time.Sleep(time.Second * 10)
		err := client.Call("Heartbeat.KeepAlive", &args, &reply)
		if err != nil {
			fmt.Println("rpc call failed:", err)
			exitChan <- err.Error()
			return
		}
	}
}

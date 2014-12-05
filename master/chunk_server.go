package master

import (
	"fmt"
	"net"
	"net/rpc"

	"GoFS/common"
)

type ChunkServer struct {
	IP   net.IP
	Port int
}

func (cs *ChunkServer) CallNewChunk(c *Chunk) error {
	addr := fmt.Sprintf("%s:%v", cs.IP.String(), cs.Port)
	client, err := rpc.Dial("tcp", addr)
	if err != nil {
		return err
	}

	args := common.NewChunkArgs{c.uuid}
	var reply common.NewChunkReply
	fmt.Println("CallNewChunk")
	err = client.Call("ChunkServer.NewChunk", &args, &reply)

	return err
}

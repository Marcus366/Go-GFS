package main

import (
	"flag"
	"fmt"
	"net"

	"GoFS/chunk"
)

var (
	MasterIP = flag.String("master-ip-address", "127.0.0.1", "IP Address of Master Server")
)

func main() {
	flag.Parse()

	ip := net.ParseIP(*MasterIP)
	if ip == nil {
		fmt.Println("Invalid Master IP Address:", *MasterIP)
		return
	}

	exitChan := make(chan string)
	c := chunk.NewChunkServer(ip)
	c.Main(exitChan)
	<-exitChan
	fmt.Println("exit")
}

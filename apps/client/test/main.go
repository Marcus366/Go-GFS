package main

import (
  "fmt"
  "log"
  "net/rpc"

  "github.com/sysu2012zzp/Go-GFS/transport"
)

func main() {
  addr := fmt.Sprintf("%s:%v", "127.0.0.1", transport.OpenClosePort)
  client, err := rpc.Dial("tcp", addr)
  if err != nil {
    log.Fatal("register fail:", err)
  }
  args := transport.OpenArgs{ "what", 1 }
  var reply transport.OpenReply
  err = client.Call("OpenClose.Open", &args, &reply)
  fmt.Println("Open:", reply.FD)
}

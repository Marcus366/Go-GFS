package main

import (
  "fmt"
  "log"
  "net/rpc"
  "time"

  "github.com/sysu2012zzp/Go-GFS/transport"
  "github.com/sysu2012zzp/Go-GFS/utils"
)

func sendHeartbeat(exitChan chan string) {
  client, err := rpc.Dial("tcp", "127.0.0.1:2345")
  if err != nil {
    log.Fatal("dialing:", err)
  }
  args := transport.HeartbeatArgs{ IP: utils.LocalIP() }
  fmt.Println("IP:", args.IP.String())
  var reply transport.HeartbeatReply
  for {
    err := client.Call("Heartbeat.KeepAlive", &args, &reply)
    if err != nil {
      fmt.Println("rpc call failed:", err)
      exitChan <- err.Error()
      return
    }
    time.Sleep(time.Second * 10)
  }
}

func main() {
  exitChan := make(chan string)
  go sendHeartbeat(exitChan)
  <-exitChan
  fmt.Println("exit")
}

package main

import (
  "fmt"
  "log"
  "net/rpc"
  "time"

  "github.com/sysu2012zzp/go-GFS/transport"
  "github.com/sysu2012zzp/go-GFS/utils"
)

func sendHeartbeat(exitChan chan string) {
  client, err := rpc.Dial("tcp", "127.0.0.1:2345")
  if err != nil {
    log.Fatal("dialing:", err)
  }
  args := transport.HeartbeatArgs{ IP: utils.LocalIP() }
  var reply transport.HeartbeatReply
  for {
    client.Go("Heartbeat.KeepAlive", &args, &reply, nil)
    /*
    <-call.Done
    if err = call.Error; err != nil {
      exitChan <- err.Error()
      fmt.Println("call error")
      return
    }
    */
    time.Sleep(time.Second * 2)
  }
}

func main() {
  exitChan := make(chan string)
  go sendHeartbeat(exitChan)
  <-exitChan
  fmt.Println("exit")
}

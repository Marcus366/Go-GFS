package main

import (
  "fmt"
  "net"

  "github.com/sysu2012zzp/Go-GFS/chunk"
)

func main() {
  exitChan := make(chan string)
  ip := net.ParseIP("127.0.0.1")
  c := chunk.NewChunk(ip)
  c.Main(exitChan)
  <-exitChan
  fmt.Println("exit")
}

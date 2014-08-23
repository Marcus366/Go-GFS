package main

import (
  "fmt"

  "github.com/sysu2012zzp/Go-GFS/gfs"
)

func main() {
  fd, err := gfs.Open("master.go")
  if err != nil {
    fmt.Println("Open fail: ", err)
    return
  }
  fmt.Println("Open succeed: ", fd)
}

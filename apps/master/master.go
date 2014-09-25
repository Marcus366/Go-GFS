package main

import (
  "fmt"
  "flag"

  "github.com/sysu2012zzp/Go-GFS/master"
  "github.com/sysu2012zzp/Go-GFS/utils"
)

var (
  version = flag.Bool("version", false, "get the version number")
)

func main() {
  flag.Parse()

  if *version {
    fmt.Println("current version: ", utils.Version())
    return
  }

  m := master.NewMaster()
  m.Main()

  select {
  }
}

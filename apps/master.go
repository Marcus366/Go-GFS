package main

import (
	"flag"
	"fmt"

	"GoFS/common"
	"GoFS/master"
)

var (
	version = flag.Bool("version", false, "get the version number")
)

func main() {
	flag.Parse()

	if *version {
		fmt.Println("current version: ", common.Version())
		return
	}

	m := master.NewMaster()
	m.Main()

	select {}
}

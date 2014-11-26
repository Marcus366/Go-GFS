package main

import (
	"flag"
	"fmt"

	"Go-GFS/gfs"
)

var (
	version = flag.Bool("version", false, "get the version number")
)

func main() {
	flag.Parse()

	if *version {
		fmt.Println("current version: ", gfs.Version())
		return
	}

	m := gfs.NewMaster()
	m.Main()

	select {}
}

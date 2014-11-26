package main

import (
	"fmt"

	"Go-GFS/gfs"
)

func main() {
	fd, err := gfs.Open("master.go")
	if err != nil {
		fmt.Println("Open fail: ", err)
		return
	}
	fmt.Println("Open succeed: ", fd)
}

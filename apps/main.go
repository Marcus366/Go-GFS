package main

import (
	"fmt"

	"GoFS/gfs"
)

func main() {
	fd, err := gfs.OpenFile("master.go", gfs.O_CREATE|gfs.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open fail: ", err)
		return
	}
	fmt.Println("Open succeed: ", fd.Fd)

	str := "what the fuck"
	_, err = fd.Write([]byte(str))
	if err != nil {
		fmt.Println("Write Error:", err)
	}

}

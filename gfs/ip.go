package gfs

import (
	"fmt"
	"net"
	"strings"
)

func LocalIP() net.IP {
	conn, err := net.Dial("udp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("localIP error:", err)
		return nil
	}
	defer conn.Close()
	return net.ParseIP(strings.Split(conn.LocalAddr().String(), ":")[0])
}

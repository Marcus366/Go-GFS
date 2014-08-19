package utils

import (
  "net"
)

func LocalIP() net.IP {
  conn, err := net.Dial("udp", "google.com:80")
  if err != nil {
    return nil
  }
  defer conn.Close()
  return net.ParseIP(conn.LocalAddr().String())
}

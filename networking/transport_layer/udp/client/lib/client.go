package lib

import (
	"fmt"
	"net"
	"os"
)

// UDPClient is a udp client
type UDPClient struct {
	DailAddr string
}

// Start will start a UDPClient instance
func (c *UDPClient) Start() {
	udpAddr, err := net.ResolveUDPAddr("udp4", c.DailAddr)
	checkError(err)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)

	_, err = conn.Write([]byte("test"))
	checkError(err)

	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)

	fmt.Println(string(buf[0:n]))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}

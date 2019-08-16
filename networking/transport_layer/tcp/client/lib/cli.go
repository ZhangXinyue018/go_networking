package lib

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// TCPClient is a client application for TCP connection
type TCPClient struct {
	DailAddr string
}

func (c *TCPClient) Conn() {
	conn, _ := net.Dial("tcp", c.DailAddr)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}

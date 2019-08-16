package lib

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// TCPServer To start a server instance
type TCPServer struct {
	Port string
}

// Start a server instance
func (s *TCPServer) Start() {
	fmt.Println("Launching server...")

	// server listener
	ln, _ := net.Listen("tcp", s.Port)

	conn, _ := ln.Accept()
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received: ", string(message))
		newMessage := strings.ToUpper(message)
		// write message to connection
		conn.Write([]byte(newMessage + "\n"))
	}

}

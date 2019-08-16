package main

import (
	"go_networking/transport_layer/udp/server/lib"
)

func main() {
	server := lib.UDPServer{
		Port: ":8081",
	}
	server.Start()
}

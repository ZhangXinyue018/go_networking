package main

import (
	"go_networking/transport_layer/tcp/server/lib"
)

func main() {
	serverEntity := lib.TCPServer{
		Port: ":8081",
	}
	serverEntity.Start()
}

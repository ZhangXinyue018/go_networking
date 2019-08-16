package main

import (
	"go_networking/transport_layer/tcp/client/lib"
)

func main() {
	clientEntity := lib.TCPClient{
		DailAddr: "localhost:8081",
	}
	clientEntity.Conn()
}

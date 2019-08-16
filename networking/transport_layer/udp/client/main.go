package main

import (
	"go_networking/transport_layer/udp/client/lib"
)

func main() {
	client := lib.UDPClient{
		DailAddr: "localhost:8081",
	}
	client.Start()
}

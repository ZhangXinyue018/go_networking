package main

import "go_networking/application_layer/ftp/client/lib"

func main() {
	client := lib.FTPClient{
		DialAddr: "localhost:8081",
	}
	client.Start()
}

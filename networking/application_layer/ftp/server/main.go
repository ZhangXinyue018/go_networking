package main

import "go_networking/application_layer/ftp/server/lib"

func main() {
	server := lib.FTPServer{
		Port: ":8081",
	}
	server.Start()
}

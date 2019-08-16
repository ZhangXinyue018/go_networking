package lib

import (
	"fmt"
	"net"
	"os"
)

const (
	DIR = "DIR"
	CD  = "CD"
	PWD = "PWD"
)

type FTPServer struct {
	Port string
}

func (fs *FTPServer) Start() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", fs.Port)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}
		go handleClient(*conn)
	}
}

func handleClient(conn net.TCPConn) {
	defer conn.Close()

	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			conn.Close()
			return
		}

		s := string(buf[0:n])
		switch {
		case s[0:2] == CD:
			chdir(conn, s[3:])
		case s[0:3] == DIR:
			dirList(conn)
		case s[0:3] == PWD:
			pwd(conn)
		}
	}
}

func chdir(conn net.TCPConn, s string) {
	err := os.Chdir(s)
	if err == nil {
		conn.Write([]byte("OK"))
	} else {
		conn.Write([]byte("ERROR"))
	}
}

func pwd(conn net.TCPConn) {
	s, err := os.Getwd()
	if err != nil {
		conn.Write([]byte(""))
		return
	}
	conn.Write([]byte(s))
}

func dirList(conn net.TCPConn) {
	defer conn.Write([]byte("\r\n"))
	dir, err := os.Open(".")
	if err != nil {
		return
	}

	names, err := dir.Readdirnames(-1)
	if err != nil {
		return
	}
	for _, nm := range names {
		conn.Write([]byte(nm + "\r\n"))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

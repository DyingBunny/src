package main

import "fmt"

import "net"

func main() {
	fmt.Println("Server start")
	ln, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("Listen error")
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept error")
			continue
		}
		go handel(conn)
	}
}

func handel(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error")
			return
		}
		fmt.Println(string(buf))
	}
}

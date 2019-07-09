package main

import (
	"fmt"
	"net"
)

type client struct {
	name string
	addr string
	pipe chan string
}

var Message = make(chan string)
var onlinemap map[string]client = make(map[string]client)

func clientToServer(client client, conn net.Conn) {
	for m := range client.pipe {
		conn.Write([]byte(m + "\n"))
	}
}

func sendMessage(name string, addr string, message string) string {
	return name + ":" + message
}

func broadcast() {
	for {
		msg := <-Message
		for _, s := range onlinemap {
			s.pipe <- msg
		}
	}
}

func handelRequest(conn net.Conn) {
	//defer conn.Close()
	addr := conn.RemoteAddr().String()
	fmt.Println(addr + "enter")
	client := client{addr, addr, make(chan string)}
	go clientToServer(client, conn)
	onlinemap[addr] = client
	Message <- sendMessage(client.name, addr, "login") //上线提示
	go func() {
		buf := make([]byte, 512)
		for {
			//fmt.Println("Please Enter Your Message:")
			n, _ := conn.Read(buf)
			if n == 0 {
				fmt.Println(client.name + "leave")
				return
			} else {
				Message <- sendMessage(client.name, addr, string(buf[:n-1]))
			}
		}
	}()
}
func main() {
	fmt.Println("Server start")
	ln, _ := net.Listen("tcp", "0.0.0.0:8888")
	defer ln.Close()
	go broadcast()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept error")
			continue
		}
		go handelRequest(conn)
	}
}

package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("connect error", err)
	}
	i := 1
	for {
		_, err := ln.Accept()
		if err != nil {
			fmt.Println("error num:", i)
			continue
		}
		fmt.Println(i, " users connected")
		i++
	}
}

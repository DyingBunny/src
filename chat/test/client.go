package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i < 50; i++ {
		_, err := net.Dial("tcp", ":8080")
		if err != nil {
			fmt.Println("Failed")
		}
		fmt.Println("Success!")
	}
}

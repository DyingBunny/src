package main

import "time"

func main() {
	for i := 1; i <= 100; i++ {
		go test_print(i)
	}
	time.Sleep(10 * time.Second)
}

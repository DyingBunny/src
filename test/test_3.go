package main

import "fmt"

func main() {
	i, n := 1, 2
	defer func(a int) {
		fmt.Println(a, n)
	}(i)
	i, n = i+1, n+2
	fmt.Println(i, n)
}

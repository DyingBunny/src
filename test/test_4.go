package main

import "fmt"

func main() {
	var input int
	var num int
	fmt.Scanln(&input)
	fmt.Scanln(&num)
	var sum int
	sum = 1
	if num == 1 {
		fmt.Println(1)
	} else {
		for i := 2; i <= num; i++ {
			sum = sum * input
		}
		fmt.Println(sum)
	}
}

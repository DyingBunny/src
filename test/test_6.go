package main

import "net/http"
import "fmt"

func main() {
	resp, _ := http.Get("https://www.baidu.com/")
	fmt.Println(resp)
}

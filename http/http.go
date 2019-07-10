package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	fmt.Fprintf(w, "hello")
}
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle login")
	fmt.Fprintf(w, "login")
}
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle history")
	fmt.Fprintf(w, "history")
}
func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/user/login", login)
	http.HandleFunc("/user/history", history)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Println("Listen failed")
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ReponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Frorm {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/home", sayhelloName)
	err := http.ListenAndServer(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}

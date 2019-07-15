package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServer("localhost:8080", nil))
}

func HandleFunc(w http.ReponseWriter, r *http.Request) {
	fmt.Printf(w, "URL.Path=%q\n", r.URL.Path)
}

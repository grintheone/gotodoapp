package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	host string = "localhost"
	port string = ":3000"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ping)
	fmt.Printf("Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(host+port, mux))
}

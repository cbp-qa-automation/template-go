package main

import (
	"flag"
	"io"
	"log"
	"net/http"
)

func main() {
	var (
		address = "127.0.0.1:8080"
	)

	flag.StringVar(&address, "address", address, "The address the server binds to.")
	flag.Parse()

	http.HandleFunc("/", helloWorld)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!\n")
}

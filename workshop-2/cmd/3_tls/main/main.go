package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Hello")
	})

	if err := http.ListenAndServeTLS(":9000", "./server.crt", "./server.key", mux); err != nil {
		log.Fatal(err)
	}
}

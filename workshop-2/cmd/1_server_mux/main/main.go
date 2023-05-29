package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":9000"
const additionalPort = ":9001"

// host: localhost:9000/
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Hello")
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("World")
	})

	go func() {
		if err := http.ListenAndServe(port, nil); err != nil {
			log.Fatal(err)
		}
	}()

	if err := http.ListenAndServe(additionalPort, mux); err != nil {
		log.Fatal(err)
	}
}

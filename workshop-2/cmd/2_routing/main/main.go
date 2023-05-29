package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":9000"

func main() {
	mux := http.NewServeMux()

	// Многоуровневый
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Hello")
	})

	// Фиксированный
	mux.HandleFunc("/article", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Article")
	})

	mux.HandleFunc("/article/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Article")
	})

	mux.HandleFunc("/article/*/name", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Regular")
	})

	mux.Handle("/home", http.RedirectHandler("http://localhost:9000/article", http.StatusPermanentRedirect))

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}

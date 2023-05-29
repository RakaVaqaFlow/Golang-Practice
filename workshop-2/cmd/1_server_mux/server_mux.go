package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	defaultMuxPort = ":9000"
	customMuxPort  = ":9001"
)

func main() {
	go func() {
		defaultMux()
	}()

	customMux()
}

func defaultMux() {
	http.HandleFunc("/", func(_ http.ResponseWriter, _ *http.Request) {
		fmt.Println("default mux")
	})

	if err := http.ListenAndServe(defaultMuxPort, nil); err != nil {
		log.Fatal(err)
	}
}

func customMux() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(_ http.ResponseWriter, _ *http.Request) {
		fmt.Println("custom mux")
	})

	if err := http.ListenAndServe(customMuxPort, mux); err != nil {
		log.Fatal(err)
	}
}

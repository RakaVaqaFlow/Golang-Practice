package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const port = ":9000"

func main() {
	var implementation server
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// Обработка HTTP-методов
		switch req.Method {
		case http.MethodGet:
			implementation.Get(res, req)
		case http.MethodDelete:
			implementation.Delete(res, req)
		case http.MethodPost:
			implementation.Create(res, req)
		case http.MethodPut:
			implementation.Update(res, req)
		default:
			fmt.Printf("unsupported method: [%s]", req.Method)
		}
	})

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

type server struct{}

func (s *server) Create(_ http.ResponseWriter, req *http.Request) {
	fmt.Printf("Create, headers: [%v]", req.Header)
}

func (s *server) Update(_ http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("error while updating, err: [%v]", err)
		return
	}

	fmt.Printf("update, body: [%s]", string(body))
}

func (s *server) Delete(_ http.ResponseWriter, _ *http.Request) {
	fmt.Println("Delete")
}

func (s *server) Get(_ http.ResponseWriter, req *http.Request) {
	fmt.Printf("Get, query params: [%v]", req.URL.Query())
}

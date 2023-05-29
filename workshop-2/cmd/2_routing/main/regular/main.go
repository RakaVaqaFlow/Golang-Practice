package main

import (
	"fmt"
	"log"
	"net/http"

	gorillamux "github.com/gorilla/mux"
)

const port = ":9000"

func main() {
	router := gorillamux.NewRouter()
	router.HandleFunc("/article/{id:[0-9]+}", func(writer http.ResponseWriter, request *http.Request) {
		vars := gorillamux.Vars(request)
		fmt.Println(vars["id"])
	})

	mux := http.NewServeMux()
	mux.Handle("/", router)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}

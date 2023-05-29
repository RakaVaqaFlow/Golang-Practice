package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port = ":9000"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/article/{category}/{id}", regularHandler)

	http.Handle("/", router)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

// ------------------------------------------------------------------------------------------------------------
// Обработчики
// ------------------------------------------------------------------------------------------------------------

func regularHandler(_ http.ResponseWriter, req *http.Request) {
	query := mux.Vars(req)
	fmt.Println(query)
}

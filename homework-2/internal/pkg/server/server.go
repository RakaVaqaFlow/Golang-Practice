package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
}

func (s *Server) Run(serverPort string) error {
	err := s.serverMux(serverPort)
	return err
}

func (s *Server) serverMux(serverPort string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.requestHandler)
	err := http.ListenAndServe(serverPort, mux)
	return err
}

func (s *Server) requestHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		s.get(res, req)
	case http.MethodPost:
		s.create(res, req)
	case http.MethodDelete:
		s.delete(res, req)
	case http.MethodPut:
		s.update(res, req)
	default:
		fmt.Printf("Unsupported method with key: [%s]\n", req.Method)
	}
}

func (s *Server) get(w http.ResponseWriter, req *http.Request) {
	queryParams := req.URL.Query()
	fmt.Printf("Get, query params: [%v]\n", queryParams)
}

func (s *Server) create(w http.ResponseWriter, req *http.Request) {
	headers := req.Header
	fmt.Printf("Create, headers:[%v]\n", headers)
	if sum := headers.Get("hw-sum"); sum != "" {
		val, err := strconv.Atoi(sum)
		if err != nil {
			fmt.Printf("hm-sum parameter should be integer value, [%s] was received\n", sum)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Printf("Sum: [%d]\n", val+5)
	}
}

func (s *Server) delete(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Delete")
}

func (s *Server) update(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error while reading request body, err: [%s]", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Printf("Update, body: [%s]\n", string(body))
}

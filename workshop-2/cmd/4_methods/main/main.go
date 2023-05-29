package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	port = ":9000"

	queryParamKey = "key"
	headerKey     = "Key"

	adminUsername = "username"
	adminPassword = "password"
)

func main() {
	implementation := server{
		data: map[string]string{},
	}

	mux := http.NewServeMux()
	mux.Handle("/", authorizationMiddleware(rootHandler(implementation)))

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}

// Authorization: Basic (<user:password> в base64)
func authorizationMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != adminUsername || password != adminPassword {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		handler.ServeHTTP(writer, r)
	}
}

func rootHandler(implementation server) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
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
	}
}

type request struct {
	Key   string
	Value string
}

type server struct {
	data map[string]string
}

func (s *server) Create(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("error while reading request body, err: [%s]", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var unmarshalled request
	if err = json.Unmarshal(body, &unmarshalled); err != nil {
		log.Printf("error while unmarshalling request body, err: [%s]", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if unmarshalled.Key == "" || unmarshalled.Value == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	if _, ok := s.data[unmarshalled.Key]; ok {
		w.WriteHeader(http.StatusConflict)
		return
	}

	s.data[unmarshalled.Key] = unmarshalled.Value
}

func (s *server) Update(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("error while reading request body, err: [%s]", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var unmarshalled request
	if err = json.Unmarshal(body, &unmarshalled); err != nil {
		log.Printf("error while unmarshalling request body, err: [%s]", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if unmarshalled.Key == "" || unmarshalled.Value == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	if _, ok := s.data[unmarshalled.Key]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	s.data[unmarshalled.Key] = unmarshalled.Value
}

func (s *server) Delete(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Query().Get(queryParamKey)
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := s.data[key]; !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	delete(s.data, key)
}

func (s *server) Get(w http.ResponseWriter, req *http.Request) {
	key := req.Header.Get(headerKey)
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	value, ok := s.data[key]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if _, err := w.Write([]byte(value)); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

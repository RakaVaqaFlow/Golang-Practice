package server_with_data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

const (
	queryParamKey = "id"
)

type ServerWithData struct {
	data map[uint32]string
}

type request struct {
	Id    uint32 `json:"id"`
	Value string `json:"value"`
}

func (s *ServerWithData) Run(serverPort string) error {
	s.data = make(map[uint32]string)
	err := s.serverMux(serverPort)
	return err
}

func (s *ServerWithData) serverMux(serverPort string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.requestHandler)
	err := http.ListenAndServe(serverPort, mux)
	return err
}

func (s *ServerWithData) requestHandler(res http.ResponseWriter, req *http.Request) {
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

func (s *ServerWithData) get(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Query().Get(queryParamKey)
	if key == "" {
		log.Print("Request parameter \"id\" is missing\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := toUint32(key)
	if err != nil {
		log.Printf("Id [%s] does not match type uint32\n", key)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	value, ok := s.data[id]
	if !ok {
		log.Printf("There is no entry in the data for this Id [%s]\n", key)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if _, err := w.Write([]byte(value)); err != nil {
		log.Printf("Error while writing value to writer, err [%s]\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *ServerWithData) create(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error while reading request body, err: [%s]\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var unmarshalled request
	if err = json.Unmarshal(body, &unmarshalled); err != nil {
		log.Printf("Error while unmarshalling request body, err: [%s]\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if unmarshalled.Value == "" {
		log.Print("Json file format error\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := s.data[unmarshalled.Id]; ok {
		log.Printf("This id=[%d] already exists", unmarshalled.Id)
		w.WriteHeader(http.StatusConflict)
		return
	}

	s.data[unmarshalled.Id] = unmarshalled.Value
}

func (s *ServerWithData) delete(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Query().Get(queryParamKey)
	if key == "" {
		log.Print("Request parameter \"id\" is missing\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := toUint32(key)
	if err != nil {
		log.Printf("Id [%s] does not match type uint32\n", key)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if _, ok := s.data[id]; !ok {
		log.Printf("There is no entry in the data for this Id [%s]\n", key)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	delete(s.data, id)
}

func (s *ServerWithData) update(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error while reading request body, err: [%s]\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var unmarshalled request
	if err = json.Unmarshal(body, &unmarshalled); err != nil {
		log.Printf("Error while unmarshalling request body, err: [%s]\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if unmarshalled.Value == "" {
		log.Print("Json file format error\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := s.data[unmarshalled.Id]; !ok {
		log.Printf("There is no entry in the data for this Id [%d]\n", unmarshalled.Id)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	s.data[unmarshalled.Id] = unmarshalled.Value
}

func toUint32(value string) (uint32, error) {
	val, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	if val < 0 {
		return 0, errors.New("value should be unsigned")
	}
	return uint32(val), nil
}

package main

import (
	. "homework/internal/pkg/server"
	. "homework/internal/pkg/server_with_data"
	"log"
)

const (
	serverPort         = ":9000"
	serverWithDataPort = ":9001"
)

func main() {

	log.Printf("Server started on localhost%s\n", serverPort)
	s1 := Server{}
	go func() {
		err := s1.Run(serverPort)
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("Server with data started on localhost%s\n", serverWithDataPort)
	s2 := ServerWithData{}
	err := s2.Run(serverWithDataPort)
	if err != nil {
		log.Fatal(err)
	}
}

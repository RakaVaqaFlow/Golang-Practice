package main

type Interface interface {
	Hello()
}

type server struct{}

func (s *server) Hello() {}

type serverMy struct{}

func (s *serverMy) Hello() {}

func main() {}

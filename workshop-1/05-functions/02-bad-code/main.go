package main

import (
	"errors"
	"fmt"
)

func doSmth() (s int, err error) {
	defer func() {
		if err != nil {
			s = 0
		}
	}()
	s = 1
	err = errors.New("test")
	return
}

func main() {
	fmt.Println(doSmth())
}

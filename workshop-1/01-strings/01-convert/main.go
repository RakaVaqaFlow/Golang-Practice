package main

import (
	"fmt"
)

func main() {
	s := "привет"
	ba := []byte(s)
	ra := []rune(s)
	fmt.Printf("%v\b\n", ba)
	fmt.Printf("%v\n\n", ra)
}

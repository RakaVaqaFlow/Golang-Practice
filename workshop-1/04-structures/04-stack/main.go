package main

import (
	"fmt"
)

type IntStack struct {
}

func (s *IntStack) Push(v int) {
}

func (s *IntStack) Pop() int {
	return 0
}

func main() {
	var s IntStack
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Printf("expected 30, got %d\n", s.Pop())
	fmt.Printf("expected 20, got %d\n", s.Pop())
	fmt.Printf("expected 10, got %d\n", s.Pop())
	fmt.Printf("expected 0, got %d\n", s.Pop())
}

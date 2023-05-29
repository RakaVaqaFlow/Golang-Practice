package main

import (
	"fmt"
)

type Base struct{}

func (b Base) Name() string {
	return "Base"
}

func (b Base) Say() {
	fmt.Println(b.Name())
}

type Child struct {
	Base
}

func (c Child) Name() string {
	return "Child"
}

func main() {
	var child Child
	child.Say() // ?
}

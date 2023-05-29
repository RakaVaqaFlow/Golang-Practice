package main

import (
	"fmt"
)

type Concater string

func (c Concater) do(i int) string {
	return fmt.Sprintf("%s: %d", c, i)
}

func main() {
	c := Concater("test")
	fmt.Println(c.do(0))

	f1 := Concater.do
	fmt.Println(f1(c, 1))

	f2 := c.do
	fmt.Println(f2(2))

	f3 := func(i int) string {
		return c.do(i)
	}
	fmt.Println(f3(3))
}

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	padding()
}

func padding() {
	fmt.Println(unsafe.Sizeof(1))   // 8 на моей машине
	fmt.Println(unsafe.Sizeof("A")) // 16 (длина + указатель)

	var x struct {
		c string // 16
		a bool   // 1
		b bool   // 1
	}

	fmt.Println(unsafe.Sizeof(x)) // 18!
	fmt.Println(
		unsafe.Offsetof(x.a), // 0
		unsafe.Offsetof(x.b), // 1
		unsafe.Offsetof(x.c)) // 8
}

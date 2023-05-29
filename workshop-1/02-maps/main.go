package main

import (
	"fmt"
)

func main() {
	m := make(map[int]struct{})
	for i := 0; i < 10; i++ {
		m[i] = struct{}{}
	}

	for i := 0; i < 10; i++ {
		//dump(m)
		//delete1If5(m)
		add12If3(m)
	}
}

func dump(m map[int]struct{}) {
	for k := range m {
		fmt.Print(k, " ")
	}
	fmt.Println()
}

func delete1If5(m map[int]struct{}) {
	for k := range m {
		if k == 5 {
			delete(m, 1)
		}
		fmt.Print(k, " ")
	}
	fmt.Println()
	m[1] = struct{}{}
}

func add12If3(m map[int]struct{}) {
	for k := range m {
		if k == 3 {
			m[12] = struct{}{}
		}
		fmt.Print(k, " ")
	}
	fmt.Println()
	delete(m, 12)
}

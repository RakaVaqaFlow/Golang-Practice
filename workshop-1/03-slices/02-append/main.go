package main

import "fmt"

func main() {
	arr := []int{1, 2}
	arr2 := arr
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println("<----------------------->")

	arr2[0] = 42
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println("<----------------------->")

	arr2 = append(arr2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println("<----------------------->")

	arr2[0] = 1
	fmt.Println(arr)
	fmt.Println(arr2)
}

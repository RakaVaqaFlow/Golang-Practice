package main

type Salary func() int

type employee struct {
	Salary Salary
}

type Boss struct {
	employee
}

type Secretary struct {
	employee
}

func main() {

}

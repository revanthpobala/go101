package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := person{"revanth", 26}
	fmt.Println("Person Name is", p.name)
	fmt.Println("Person Age is", p.age)
}

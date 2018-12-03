package main

import "fmt"

func main() {

	i:= 7
	fmt.Println("Initial Address ",&i, "initial value", i)
	inc(&i)
	fmt.Println("Address after increment",&i, "value", i)
}

func inc(x *int)  {
	*x++
}
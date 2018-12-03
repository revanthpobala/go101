package main

import "fmt"

func main() {

	//Maps

	vertices := make(map[string]string)

	vertices["s"] = "sujatha"
	vertices["r"] =  "Rapolu"

	fmt.Println("The vertices are " , vertices)

	//Lists

	lists := []int {4,5,6,8}
	d := append(lists, 100)
	fmt.Println(d)

}

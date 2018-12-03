package main

func main() {
	//For loop

	for i := 0; i < 10; i++ {
		println(i)
	}

	// While loop

	var arr = []string{"a", "b", "c"}

	for index, value := range arr {
		println("index", index, value)
	}
}

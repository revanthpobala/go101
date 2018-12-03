package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 7

	if x > 5 {
		fmt.Println("Greater than 5")
	} else if x < 5 {
		fmt.Println("Less than 5")
	} else {
		fmt.Println("The number you entered is " + strconv.Itoa(x))
	}

}

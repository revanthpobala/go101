package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	var result, err = fmt.Println(sqrt(16 ))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

 }

func sqrt( x float64) (float64, error)  {

	if x < 0 {
		//noinspection GoTypesCompatibility
		return 0, errors.New("values cannot be less than 0")
	}

	return math.Sqrt(x), nil
}

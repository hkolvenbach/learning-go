package main

import (
	"errors"
	"fmt"
)

func main() {
	a := 12
	b := 0

	// c := divideTwoNumbers(a, b)

	// if c == 2 {
	// 	fmt.Println("Found two")
	// }

	c, err := divideTwoNumbers(a, b)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		if c == 2 {
			fmt.Println("Found a two")
		}
	}

	//fmt.Println()
}

func divideTwoNumbers(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return x / y, nil
}

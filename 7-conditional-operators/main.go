package main

import "fmt"

func main() {
	a := 12
	b := 0

	// c := divideTwoNumbers(a, b)

	// if c == 2 {
	// 	fmt.Println("Found two")
	// }

	if b != 0 && divideTwoNumbers(a, b) == 2 {
		fmt.Println("Found a two")
	}

	//fmt.Println()
}

func divideTwoNumbers(x, y int) int {
	return x / y
}

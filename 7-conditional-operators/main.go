package main

import "fmt"

func main() {
	a := 12
	b := 6

	c := divideTwoNumbers(a, b)

	if c == 2 {
		fmt.Println("Found two")
	}
	//fmt.Println()
}

func divideTwoNumbers(x, y int) int {
	return x / y
}

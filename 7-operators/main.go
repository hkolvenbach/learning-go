package main

import (
	"fmt"
	"math"
)

func main() {
	// area = pi*r^2
	var radius = 12.0
	area := math.Pi * radius * radius
	fmt.Println("Area", area)

	// division
	var half float32 = 1.0 / 2.0
	fmt.Println("half with integer division", half)

	// 3 squared
	squared := math.Pow(3.0, 2.0)
	fmt.Println("3^2=", squared)

}

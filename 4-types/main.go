package main

import "fmt"

func main() {
	var myInt int = 10

	myFirstPointer := &myInt

	fmt.Println(myInt)
	fmt.Println("myFirstPointer is", myFirstPointer)

	*myFirstPointer = 15
	fmt.Println(myInt)
}

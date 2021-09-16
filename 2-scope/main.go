package main

import (
	"fmt"
	"myapp/packageone"
)

var one = "One"

func main() {
	var one = "this is a block level variable"
	fmt.Println(one)

	myFunc()

	newString := packageone.PublicVar

	fmt.Println(newString)
}

func myFunc() {
	//var one = "the number one"
	fmt.Println(one)
}

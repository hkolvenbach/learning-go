package main

import "fmt"

func main() {
	age := 10
	name := "Jack"
	rightHanded := true

	fmt.Printf("%s is %d years old. Right handed: %t", name, age, rightHanded)
	fmt.Println()

	// expressions
	age = age + 10
	isTeenager := age >= 13
	fmt.Printf("Is teenager: %t\n", isTeenager)
}

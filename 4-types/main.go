package main

import "fmt"

type Animal struct {
	Name  string
	Sound string
}

func main() {
	myAnimal := Animal{Name: "dog", Sound: "woof"}
	myAnimal.Says()
}

func (a *Animal) Says() {
	fmt.Println("Animal", a.Name, "says", a.Sound)
}

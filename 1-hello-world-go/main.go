package main

import "fmt"

func main()  {
	var whatToSay string = "Hello world variable."
	sayHelloWorld(whatToSay)
}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
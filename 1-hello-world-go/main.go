package main

import "fmt"

func main()  {
	sayHelloWorld("Hello World, again.")
}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
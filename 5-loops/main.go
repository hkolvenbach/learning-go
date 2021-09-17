package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	i := 1000
	for i > 100 {
		// get random number
		i = rand.Intn(1000) + 1
		fmt.Println("I is", i)
		if i > 100 {
			fmt.Println("is is", i, "so loop keeps going")
		}
	}
	fmt.Println("Got", i, "and broke out of the loop")
}

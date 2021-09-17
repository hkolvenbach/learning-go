package main

import "fmt"

func main() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "cat")

	fmt.Println(animals)

	for _, x := range animals {
		fmt.Println(x)
	}
}

func DeleteFromSlice(a []string, i int) {
	// copy last element to ith element (which will be deleted)
	// this will not keep the order
	a[i] = a[len(a)-1]
	// set last element to nothing
	a[len(a)-1] = ""
	// use 0 to len(a)-1 elements and assign to a
	a = a[:len(a)-1]
}

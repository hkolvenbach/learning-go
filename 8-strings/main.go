package main

import "fmt"

func main() {
	fmt.Println()
	name := "Hello World"
	fmt.Println("String", name)

	fmt.Println("Bytes")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x", name[i])
	}

	fmt.Println()
	fmt.Println()

	fmt.Println("Index\tRune\tString")

	for x, y := range name {
		fmt.Println(x, "\t", y, "\t", string(y))
		//fmt.Printf("%d\t%X\t%s", i, name[i], name[i])
	}

}

package main

import (
	"fmt"
	"strings"
)

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

	h := "Hello "
	w := "World"
	myString := h + w
	fmt.Println(myString)
	myString = fmt.Sprintf("%s%s", h, w)
	fmt.Println(myString)

	var sb strings.Builder
	sb.WriteString(h)
	sb.WriteString(w)
	fmt.Println(sb.String())

	fmt.Println()
	name = "abcdefghijkl"
	fmt.Println("Getting Substring")
	fmt.Println(name[0:10])
}

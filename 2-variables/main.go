package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "and press ENTER when ready."

func main() {
	// declare, than assign
	var firstNumber = 2
	var secondNumber = 5
	var substraction = 7
	// var answer int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guess theNumbers Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// Inserts space automatically between arguments
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("No multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally though of", prompt)
	reader.ReadString('\n')

	fmt.Println("no substract", substraction, prompt)
	reader.ReadString('\n')

}

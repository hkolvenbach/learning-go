package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappuchino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("----")
	for i := 1; i <= len(coffees); i++ {
		fmt.Printf("%d - %s\n", i, coffees[i])
	}

	char := ' '

	for char != 'q' {
		char, _, err = keyboard.GetSingleKey()
		if char == 'Q' {
			char = 'q'
		}
		if err != nil {
			log.Fatal(err)
		}

		// if char == 'q' || char == 'Q' {
		// 	break
		// }

		if char != 'q' {
			i, _ := strconv.Atoi(string(char))

			fmt.Printf("You chose %s\n", coffees[i])
		}

	}

	fmt.Println("Thanks for visiting us. Have a great day.")
	fmt.Println("Program exiting.")
}

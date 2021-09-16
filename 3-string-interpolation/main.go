package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

type User struct {
	UserName        string
	Age             int
	FavouriteNumber float64
	OwnsADog        bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	user.UserName = readString("What is your name?")

	user.Age = readInt("How old are you?")

	user.FavouriteNumber = readFloat64("What is your favorite number?")

	user.OwnsADog = readBool("Do you own a dog, y/n?")

	fmt.Printf("Your name is %s. You are %d years old. Your favorite number is %.2f. ", user.UserName, user.Age, user.FavouriteNumber)
	if user.OwnsADog {
		fmt.Println("And you own a dog.")
	} else {
		fmt.Println("And you do not own a dog.")
	}
}

func prompt() {
	fmt.Print("->")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "" {
			fmt.Println("Please enter a value")
		} else {
			return userInput
		}
	}
}

func readInt(s string) int {

	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}

}

func readFloat64(s string) float64 {

	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}

}

func readBool(s string) bool {
	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(s)
		prompt()
		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}

		// alternative:
		// strings.toLower(string(char)) != 'y' && strings.toLower(string(char)) != 'n'
		if char == 'y' || char == 'Y' {
			return true
		} else if char == 'n' || char == 'N' {
			return false
		} else {
			fmt.Println("Please enter 'y' or 'n'")
		}
	}

}

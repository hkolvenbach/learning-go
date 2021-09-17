package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

var keyPressChannel chan rune

func main() {
	keyPressChannel = make(chan rune)
	go listenForKeyPress()

	fmt.Println("Press any key, q to quit")
	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}

		keyPressChannel <- char
	}
}

// this is a blocking go routine (blocks until receiving something from channel)
func listenForKeyPress() {
	for {
		key := <-keyPressChannel
		fmt.Println("You pressed", string(key))
	}
}

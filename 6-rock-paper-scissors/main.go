package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {
	rand.Seed(time.Now().UnixNano())
	playerChoice := ""
	playerValue := -1

	computerValue := rand.Intn(2)

	reader := bufio.NewReader(os.Stdin)

	clearScreen()

	for i := 1; i <= 3; i++ {
		fmt.Printf("Round %d:\n", i)
		fmt.Print("Please enter rock, paper, or scissors ->")
		playerChoice, _ = reader.ReadString('\n')
		playerChoice = strings.Replace(playerChoice, "\r\n", "", -1)
		playerChoice = strings.Replace(playerChoice, "\n", "", -1)

		if playerChoice == "rock" {
			playerValue = ROCK
		} else if playerChoice == "paper" {
			playerValue = PAPER
		} else if playerChoice == "scissors" {
			playerValue = SCISSORS
		} else {
			// reset if invalid data is entered
			playerValue = -1
		}

		switch computerValue {
		case ROCK:
			fmt.Println("Computer chose ROCK")
			break
		case PAPER:
			fmt.Println("Computer chose PAPER")
			break
		case SCISSORS:
			fmt.Println("Computer chose SCISSORS")
			break
		default:
		}
		fmt.Println("Player value is", playerValue)

		if playerValue == computerValue {
			fmt.Println("Its a draw.")
		} else {
			switch playerValue {
			case ROCK:
				if computerValue == PAPER {
					fmt.Println("Computer wins")
				} else {
					fmt.Println("Player wins")
				}
				break
			case PAPER:
				if computerValue == SCISSORS {
					fmt.Println("Computer wins")
				} else {
					fmt.Println("Player wins")
				}
				break
			case SCISSORS:
				if computerValue == ROCK {
					fmt.Println("Computer wins")
				} else {
					fmt.Println("Player wins")
				}
				break
			default:
				fmt.Println("Invalid choice")
				// reset round counter by one
				i--
			}
		}
	}

}

// clearScreen clears the screen
func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

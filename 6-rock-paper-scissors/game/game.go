package game

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

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Round       Round
}

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

var reader = bufio.NewReader(os.Stdin)

func (g *Game) Rounds() {
	// use select to process input into channels
	// Print to screen
	// keep track of round number
	for {
		select {
		case round := <-g.RoundChan:
			g.Round.RoundNumber = g.Round.RoundNumber + round
			g.RoundChan <- 1
		case msg := <-g.DisplayChan:
			fmt.Println(msg)
			// Notify that it is finished
			g.DisplayChan <- ""
		}
	}
}

// clearScreen clears the screen
func (g *Game) ClearScreen() {
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

func (g *Game) PrintIntro() {
	g.DisplayChan <- "Please enter rock, paper, or scissors"
	// wait for it to finish printing
	<-g.DisplayChan
}

func (g *Game) PlayRound() bool {
	rand.Seed(time.Now().UnixNano())
	playerValue := -1
	g.DisplayChan <- ""
	// wait for it to finish printing
	<-g.DisplayChan
	g.DisplayChan <- fmt.Sprintf("Round %d", g.Round.RoundNumber)

	// wait for it to finish printing
	<-g.DisplayChan

	fmt.Print("Please enter rock, paper, or scissors ->")
	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\r\n", "", -1)
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	computerValue := rand.Intn(2)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	fmt.Println()
	g.DisplayChan <- fmt.Sprintf("Player chose %s", playerChoice)
	// wait for it to finish printing
	<-g.DisplayChan
	g.DisplayChan <- fmt.Sprintf("Player chose %s", strings.ToUpper(playerChoice))
	// wait for it to finish printing
	<-g.DisplayChan

	switch computerValue {
	case ROCK:
		g.DisplayChan <- "Computer chose ROCK"
		// wait for it to finish printing
		<-g.DisplayChan
		break
	case PAPER:
		g.DisplayChan <- "Computer chose PAPER"
		// wait for it to finish printing
		<-g.DisplayChan
		break
	case SCISSORS:
		g.DisplayChan <- "Computer chose SCISSORS"
		// wait for it to finish printing
		<-g.DisplayChan
		break
	default:
	}

	if playerValue == computerValue {
		g.DisplayChan <- "Its a draw."
		// wait for it to finish printing
		<-g.DisplayChan
		return false
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		case PAPER:
			if computerValue == SCISSORS {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		case SCISSORS:
			if computerValue == ROCK {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		default:
			g.DisplayChan <- "Invalid choice"
			// wait for it to finish printing
			<-g.DisplayChan
			// reset round counter by one
			return false
		}
	}
	return true
}

func (g *Game) computerWins() {
	g.Round.ComputerScore++
	g.DisplayChan <- "Computer wins!"
	// wait for it to finish printing
	<-g.DisplayChan
}

func (g *Game) playerWins() {
	g.Round.PlayerScore++
	g.DisplayChan <- "Player wins!"
	// wait for it to finish printing
	<-g.DisplayChan
}

func (g *Game) PrintSummary() {
	g.DisplayChan <- ""
	// wait for it to finish printing
	<-g.DisplayChan
	g.DisplayChan <- fmt.Sprintf("Computer Score: %d", g.Round.ComputerScore)
	// wait for it to finish printing
	<-g.DisplayChan
	g.DisplayChan <- fmt.Sprintf("Player Score: %d", g.Round.PlayerScore)
	// wait for it to finish printing
	<-g.DisplayChan
	if g.Round.ComputerScore == g.Round.PlayerScore {
		g.DisplayChan <- "It's a draw!"
		// wait for it to finish printing
		<-g.DisplayChan
	} else if g.Round.ComputerScore > g.Round.PlayerScore {
		g.DisplayChan <- "Computer wins!"
		// wait for it to finish printing
		<-g.DisplayChan
	} else {
		g.DisplayChan <- "Player wins!"
		// wait for it to finish printing
		<-g.DisplayChan
	}
}

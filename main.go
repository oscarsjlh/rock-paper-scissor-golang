package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	options := []string{"rock", "paper", "scissor"}
	// chosen := selector(options),
	fmt.Println("Rock paper or scissor? (Enter q to exit)")
	for true {
		userInput := choice()
		if userInput == "q" {
			println("Sad to see you go")
			os.Exit(0)
		}
		for !contains(options, userInput) {
			fmt.Println("Input not valid try again")
			userInput = choice()
		}

		fmt.Println("Valid input.")
		m := selector(options)
		res := whowon(m, userInput)
		println(res)
		println("Wanna play again (y/n)")
		playAgain := choice()
		if playAgain != "y" {
			break
		}
	}
}

func selector(possibilities []string) string {
	rand.Seed(time.Now().UnixNano())
	pick := possibilities[rand.Intn(len(possibilities))]

	return pick
}

func choice() string {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	trimedstr := strings.TrimSpace(str)
	return trimedstr
}

func contains(s []string, str string) bool {
	for _, b := range s {
		if b == str {
			return true
		}
	}
	return false
}

func whowon(machine string, user string) string {
	var result string

	if machine == user {
		result = "Tie"
	} else if user == "paper" {
		if machine == "rock" {
			result = "Human wins"
		} else {
			result = "Machine wins"
		}
	} else if user == "rock" {
		if machine == "paper" {
			result = "machine wins"
		} else {
			result = "User wins"
		}
	} else if user == "scissor" {
		if machine == "paper" {
			result = "user wins"
		} else {
			result = "machine wins"
		}
	}

	return result
}

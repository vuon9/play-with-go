package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	println("Guess the number!")

	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(101)

	for {
		fmt.Printf("Please input your guess: ")

		var err error
		var guess int
		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			guess, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Printf("Your input is not a number!")
				continue
			}
		}

		if guess > secretNumber {
			fmt.Println("Too big!")
		}

		if guess < secretNumber {
			fmt.Println("Too small!")
		}

		if guess == secretNumber {
			fmt.Println("You win!")
			break
		}

	}
}

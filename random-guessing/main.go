package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	secretNum := RandInt(1, 101)

	fmt.Println("I'm thinking of a number between 1 and 100")

	for {
		var guess int
		fmt.Print("What's your guess? ")
		fmt.Scanf("%d\n", &guess)

		if guess > secretNum {
			fmt.Println("Too high")
		} else if guess < secretNum {
			fmt.Println("Too low")
		} else {
			fmt.Println("Correct! You win!")
			break
		}
	}
}

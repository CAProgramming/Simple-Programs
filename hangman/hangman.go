package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func Contains(arr []rune, c rune) bool {
	for _, value := range arr {
		if value == c {
			return true
		}
	}
	return false
}

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	text, err := ioutil.ReadFile("words.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	words := strings.Split(string(text), "\n")
	secretWord := words[RandInt(0, len(words))]

	fmt.Println("Welcome to hangman!")

	var dashes []rune

	for range secretWord {
		dashes = append(dashes, '-')
	}

	for {
		fmt.Println(string(dashes))

		var guess rune

		fmt.Print("What's your guess? ")
		fmt.Scanf("%c\n", &guess)

		for i, c := range secretWord {
			if c == guess {
				dashes[i] = c
			}
		}

		if !Contains(dashes, '-') {
			fmt.Println("You win!")
			fmt.Println("The word was", secretWord)
			break
		}
	}
}

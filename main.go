package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	word := words[rand.Intn(len(words))]
	lives := len(word)

	var blanks []string
	for range word {
		blanks = append(blanks, "_")
	}

	var guess string
	for {
		fmt.Printf("❤️ %d Lives, Word: %s Guess:", lives, strings.Join(blanks, ""))
		fmt.Scan(&guess)

		for _, input := range guess {
			correct := false

			for i, letter := range word {
				if letter == input {
					blanks[i] = string(letter)
					correct = true
				}
			}

			if !correct {
				lives--
			}
		}

		if lives <= 0 {
			fmt.Printf("💔 You lost! The word was %s.\n", word)
			break
		}

		if word == strings.Join(blanks, "") {
			fmt.Println("⭐️ You won! Great job!")
			break
		}
	}
}

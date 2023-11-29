package main

import (
	"bufio"
	"os"
)

var words = []string{}

func loadWords() {
	file, _ := os.Open("words.txt")
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		words = append(words, sc.Text())
	}
}

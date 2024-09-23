package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	rand.Seed(time.Now().UnixNano())
	RandomWord := lines[rand.Intn(len(lines))]

	word := RandomWord
	fmt.Println(word)
	var i int = 9
	var test string
	fmt.Println("Enter the word: ")
	fmt.Scanln(&test)
	for i > 0 {
		if test == RandomWord {
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect!", i, "attempts left")
			i--
		}
	}
}

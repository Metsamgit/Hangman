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

	var RandomLetters []int
	for len(RandomLetters) < 3 {
		r := rand.Intn(len(RandomWord))
		alreadyChosen := false
		for _, v := range RandomLetters {
			if v == r {
				alreadyChosen = true
				break
			}
		}
		if !alreadyChosen {
			RandomLetters = append(RandomLetters, r)
		}
	}
	word := make([]rune, len(RandomWord))
	for i := range word {
		word[i] = '_'
	}
	for _, v := range RandomLetters {
		word[v] = rune(RandomWord[v])
	}
	fmt.Println("Word to guess:", string(word))

	var i int = 10
	var test string
	fmt.Println("Enter the word: ")
	fmt.Scanln(&test)
	for i > 0 {
		if test == RandomWord {
			fmt.Println("Correct!")
			break
		} else {
			fmt.Println("Incorrect!", i, "attempts left")
			i--
			fmt.Println("Enter the word: ")
			fmt.Scanln(&test)
		}
	}
}

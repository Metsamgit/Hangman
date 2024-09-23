package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	wordFile, err := os.Open("words.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wordFile.Close()

	var lines []string
	scanner := bufio.NewScanner(wordFile)
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

	var attempts int = 9
	var guess string
	fmt.Println("Enter the word: ")
	fmt.Scanln(&guess)

	stages := readHangmanStages("hangmanposition.txt")

	for attempts > 0 {
		if guess == RandomWord {
			fmt.Println("Correct!")
			break
		} else {
			fmt.Println("Incorrect!", attempts, "attempts left")
			attempts--

			if attempts < len(stages) {
				fmt.Println(stages[len(stages)-attempts-1])
			}

			fmt.Println("Enter the word: ")
			fmt.Scanln(&guess)
		}
	}

	if attempts == 0 {
		fmt.Println("Game over! The word was:", RandomWord)
	}
}

func readHangmanStages(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading hangman stage file:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var stages []string
	var currentStage strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {

			stages = append(stages, currentStage.String())
			currentStage.Reset()
		} else {

			currentStage.WriteString(line + "\n")
		}
	}

	if currentStage.Len() > 0 {
		stages = append(stages, currentStage.String())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}

	return stages
}

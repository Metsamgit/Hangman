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
	// Ouvrir le fichier de mots
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

	// Sélectionner un mot aléatoire
	rand.Seed(time.Now().UnixNano())
	RandomWord := lines[rand.Intn(len(lines))]
	RandomWord = strings.ToLower(RandomWord) // Normaliser les entrées en minuscules

	// Créer un tableau pour afficher le mot avec des underscores pour tout sauf les 3 dernières lettres
	word := make([]rune, len(RandomWord))
	for i := range word {
		// Révéler les 3 dernières lettres du mot
		if i >= len(RandomWord)-3 {
			word[i] = rune(RandomWord[i])
		} else {
			word[i] = '_'
		}
	}

	// Afficher le mot avec les 3 dernières lettres découvertes
	fmt.Println("Word to guess:", string(word))

	// Gérer les tentatives
	attempts := 9

	// Lire et afficher les étapes du pendu
	stages := readHangmanStages("hangmanposition.txt")

	// Boucle de jeu : demander à l'utilisateur de deviner des lettres
	for attempts > 0 {
		var guess string
		fmt.Print("Enter a letter: ")
		fmt.Scanln(&guess)
		guess = strings.ToLower(guess) // Normaliser l'entrée

		if len(guess) != 1 {
			fmt.Println("Please enter a single letter.")
			continue
		}

		correctGuess := false
		for i := 0; i < len(RandomWord)-3; i++ { // Ne pas permettre de deviner les 3 dernières lettres
			if rune(guess[0]) == rune(RandomWord[i]) && word[i] == '_' {
				word[i] = rune(RandomWord[i])
				correctGuess = true
			}
		}

		if correctGuess {
			fmt.Println("Good guess!", string(word))
		} else {
			fmt.Println("Incorrect guess!", attempts, "attempts left")
			attempts--

			// Afficher l'étape du pendu correspondante
			if attempts < len(stages) {
				fmt.Println(stages[len(stages)-attempts-1])
			}
		}

		// Vérifier si le mot est entièrement découvert
		if string(word) == RandomWord {
			fmt.Println("Congratulations! You've guessed the word:", RandomWord)
			return
		}
	}

	// Si toutes les tentatives sont épuisées
	fmt.Println("Game over! The word was:", RandomWord)
}

// Fonction pour lire les étapes du pendu à partir d'un fichier
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
			// Ajouter l'étape complète au tableau et réinitialiser pour la prochaine étape
			stages = append(stages, currentStage.String())
			currentStage.Reset()
		} else {
			// Construire l'étape actuelle
			currentStage.WriteString(line + "\n")
		}
	}

	// Ajouter la dernière étape s'il n'y a pas de ligne vide à la fin du fichier
	if currentStage.Len() > 0 {
		stages = append(stages, currentStage.String())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}

	return stages
}

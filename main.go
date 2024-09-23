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
	fichierMots, err := os.Open("words.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fichierMots.Close()

	var lignes []string
	scanner := bufio.NewScanner(fichierMots)
	for scanner.Scan() {
		lignes = append(lignes, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	rand.Seed(time.Now().UnixNano())
	motAleatoire := lignes[rand.Intn(len(lignes))]
	motAleatoire = strings.ToLower(motAleatoire)

	mot := make([]rune, len(motAleatoire))
	for i := range mot {
		if i >= len(motAleatoire)-3 {
			mot[i] = rune(motAleatoire[i])
		} else {
			mot[i] = '_'
		}
	}

	fmt.Println("Mot à deviner :", string(mot))

	tentatives := 9
	etapes := lireEtapesPendu("hangmanposition.txt")

	for tentatives > 0 {
		var proposition string
		fmt.Print("Entrez une lettre ou devinez le mot complet : ")
		fmt.Scanln(&proposition)
		proposition = strings.ToLower(proposition)

		if len(proposition) > 1 {
			if proposition == motAleatoire {
				fmt.Println("Félicitations ! Vous avez deviné le mot :", motAleatoire)
				return
			} else {
				fmt.Println("Mot incorrect !", tentatives, "tentatives restantes")
				tentatives--
				if tentatives < len(etapes) {
					fmt.Println(etapes[len(etapes)-tentatives-1])
				}
			}
		} else {
			bonneProposition := false
			for i := 0; i < len(motAleatoire)-3; i++ {
				if rune(proposition[0]) == rune(motAleatoire[i]) && mot[i] == '_' {
					mot[i] = rune(motAleatoire[i])
					bonneProposition = true
				}
			}

			if bonneProposition {
				fmt.Println("Bonne lettre !", string(mot))
			} else {
				fmt.Println("Lettre incorrecte !", tentatives, "tentatives restantes")
				tentatives--
				if tentatives < len(etapes) {
					fmt.Println(etapes[len(etapes)-tentatives-1])
				}
			}

			if string(mot) == motAleatoire {
				fmt.Println("Félicitations ! Vous avez deviné le mot :", motAleatoire)
				return
			}
		}
	}

	fmt.Println("Fin de partie ! Le mot était :", motAleatoire)
}

func lireEtapesPendu(nomFichier string) []string {
	fichier, err := os.Open(nomFichier)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return nil
	}
	defer fichier.Close()

	scanner := bufio.NewScanner(fichier)
	var etapes []string
	var etapeActuelle strings.Builder

	for scanner.Scan() {
		ligne := scanner.Text()
		if ligne == "" {
			etapes = append(etapes, etapeActuelle.String())
			etapeActuelle.Reset()
		} else {
			etapeActuelle.WriteString(ligne + "\n")
		}
	}

	if etapeActuelle.Len() > 0 {
		etapes = append(etapes, etapeActuelle.String())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur de scanner :", err)
	}

	return etapes
}

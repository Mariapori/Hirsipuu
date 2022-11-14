package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	welcomeText := "Tervetuloa pelaamaan hirsipuuta"
	fmt.Println(welcomeText)
	words := []string{"koskenkorva", "koskenlaskija", "marlboro", "torvalds", "mariapori", "digiruokalista", "helsinki", "tampere", "rovaniemi", "lentokone"}
	for i := 0; i < len(words); i++ {
		word := words[i]
		fmt.Print(randomizeWord(word) + "\n")
	}
}

func randomizeWord(word string) string {
	mutWord := word
	rand.Seed(time.Now().UnixNano())
	wordLength := rand.Intn(len(word))
	for i := 0; i < wordLength/3; i++ {
		char := string([]rune(word)[i])
		mutWord = strings.Replace(mutWord, char, "_", -1)
	}
	return mutWord
}

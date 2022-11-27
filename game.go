package main

import "math/rand"

func gameHandler() {
	word := getWord()
	LoggerInfo.Printf("Game word: %v\n", word)
}

func getWord() string {
	// static word list for now
	wordList := []string{"audio", "block", "chain", "flake",
		"lunch", "noise", "price", "stare", "teach", "zebra"}
	randomIndex := rand.Intn(len(wordList))
	return wordList[randomIndex]
}

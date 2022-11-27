package main

import (
	"fmt"
	"math/rand"
)

var lettersGray []string
var lettersGreen []string
var lettersYellow []string
var lettersInWord [5]string

func gameHandler() {
	populateGrayLetters()
	populateLettersInWord()
	result := false
	word := getWord()
	if debugFlow {
		LoggerDebug.Printf("Game word: %v\n", word)
	}
	for attempt := 1; attempt <= 6; attempt++ {
		renderTerminal(attempt)
		userWord := userInput()
		result = analyzeWord(userWord, word)
		if result {
			break
		}
	}
	renderTerminalFinal(result, word)

}

func analyzeWord(guessed string, actual string) bool {
	if guessed == actual {
		fmt.Printf("Guessed: %s, Actual: %s\n", guessed, actual)
		return true
	}
	return false
}

func getWord() string {
	// static word list for now
	wordList := []string{"audio", "block", "chain", "flake",
		"lunch", "noise", "price", "stare", "teach", "zebra"}
	randomIndex := rand.Intn(len(wordList))
	return wordList[randomIndex]
}

func populateGrayLetters() {
	// there probably is a more efficient way
	lettersGray = append(lettersGray, "a")
	lettersGray = append(lettersGray, "b")
	lettersGray = append(lettersGray, "c")
	lettersGray = append(lettersGray, "d")
	lettersGray = append(lettersGray, "e")
	lettersGray = append(lettersGray, "f")
	lettersGray = append(lettersGray, "g")
	lettersGray = append(lettersGray, "h")
	lettersGray = append(lettersGray, "i")
	lettersGray = append(lettersGray, "j")
	lettersGray = append(lettersGray, "k")
	lettersGray = append(lettersGray, "l")
	lettersGray = append(lettersGray, "m")
	lettersGray = append(lettersGray, "n")
	lettersGray = append(lettersGray, "o")
	lettersGray = append(lettersGray, "p")
	lettersGray = append(lettersGray, "q")
	lettersGray = append(lettersGray, "r")
	lettersGray = append(lettersGray, "s")
	lettersGray = append(lettersGray, "t")
	lettersGray = append(lettersGray, "u")
	lettersGray = append(lettersGray, "v")
	lettersGray = append(lettersGray, "w")
	lettersGray = append(lettersGray, "x")
	lettersGray = append(lettersGray, "y")
	lettersGray = append(lettersGray, "z")
}

func populateLettersInWord() {
	lettersInWord[0] = "_"
	lettersInWord[1] = "_"
	lettersInWord[2] = "_"
	lettersInWord[3] = "_"
	lettersInWord[4] = "_"
}

func renderTerminal(attempt int) {
	fmt.Printf("\n\n")
	fmt.Printf("On try %d out of 6\n", attempt)
	fmt.Printf("Gray letters: %v\n", lettersGray)
	fmt.Printf("Green letters: %v\n", lettersGreen)
	fmt.Printf("Yellow letters: %v\n", lettersYellow)
	fmt.Printf("Wordle word: %v\n", lettersInWord)
	fmt.Printf("\n\n")
}

func renderTerminalFinal(game bool, word string) {
	fmt.Printf("\n\n")
	if game {
		fmt.Println("You won!!!")
	} else {
		fmt.Printf("Word was: %s\n", word)
		fmt.Println("Better luck next time.")
	}
	fmt.Printf("\n\n")
}

func userInput() string {
	fmt.Println("Word guess: ")
	var input string
	fmt.Scanln(&input)
	valid, message := validateWord(input)
	if valid {
		fmt.Println(message)
	} else {
		fmt.Println(message)
		userInput()
	}
	// fmt.Print(input)
	fmt.Printf("\n")
	return input
}

func validateWord(word string) (bool, string) {
	if len(word) != 5 {
		message := "Word does not have 5 letters"
		// fmt.Println(message)
		return false, message
	}
	// TODO: validate if actual dictionary word
	return true, "Word is valid"
}

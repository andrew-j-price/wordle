package main

import (
	"fmt"
	"math/rand"
	"sort"
)

var lettersAvailable []string
var lettersGray []string
var lettersGreen []string
var lettersYellow []string
var lettersInWord [5]string
var sliceOfWord []string

func gameHandler() {
	populateAvailableLetters()
	populateLettersInWord() // defaults to _ _ _ _ _
	result := false         // defaults to game loss
	word := getWord()       // gets randomly generated word
	sliceOfWord = sliceOfString(word)
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
	guessSlice := sliceOfString(guessed)
	actualSlice := sliceOfString(actual)
	// fmt.Printf("guessSlice: is of type: %v, with value: %v\n", reflect.TypeOf(guessSlice), guessSlice)
	// fmt.Printf("actualSlice: is of type: %v, with value: %v\n", reflect.TypeOf(actualSlice), actualSlice)
	for i := 0; i < 5; i++ {
		analyzeLetter(guessSlice[i], actualSlice[i], actual, i)
	}
	for _, r := range lettersYellow {
		// fmt.Printf("Analyzing yellow letter: %s", r)
		if stringInSlice(r, lettersGreen) {
			lettersYellow = removeStringFromSlice(lettersYellow, r)
		}
	}
	if guessed == actual {
		fmt.Printf("Guessed: %s, Actual: %s\n", guessed, actual)
		return true
	}
	return false
}

func analyzeLetter(guessLetter string, actualLetter string, actualWord string, index int) {
	// fmt.Printf("Analyzing if: %s, matches: %s, and in: %s\n", guessLetter, actualLetter, actualWord)
	// If Letter Matches
	if guessLetter == actualLetter {
		if !(stringInSlice(guessLetter, lettersGreen)) {
			lettersGreen = append(lettersGreen, guessLetter)
		}
		lettersInWord[index] = guessLetter
	}

	// If Letter is not in word
	// TODO: handle double letters
	// TODO: reference slice of actual word globally
	if !(stringInSlice(guessLetter, sliceOfString(actualWord))) {
		if !(stringInSlice(guessLetter, lettersGray)) {
			lettersGray = append(lettersGray, guessLetter)
		}
	}

	// If Letter is in the word
	if stringInSlice(guessLetter, sliceOfString(actualWord)) {
		if !(stringInSlice(guessLetter, lettersYellow)) && !(stringInSlice(guessLetter, lettersGreen)) {
			lettersYellow = append(lettersYellow, guessLetter)
		}
	}

	// Remove letter from the guessed letters
	if stringInSlice(guessLetter, lettersAvailable) {
		lettersAvailable = removeStringFromSlice(lettersAvailable, guessLetter)
	}
}

func getWord() string {
	// static word list for now
	wordList := []string{"audio", "block", "chain", "flake",
		"lunch", "noise", "price", "stare", "teach", "zebra"}
	randomIndex := rand.Intn(len(wordList))
	return wordList[randomIndex]
}

func populateAvailableLetters() {
	// there probably is a more efficient way
	lettersAvailable = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
}

func populateLettersInWord() {
	lettersInWord = [5]string{"_", "_", "_", "_", "_"}
}

func renderTerminal(attempt int) {
	sort.Strings(lettersGray)
	sort.Strings(lettersGreen)
	sort.Strings(lettersYellow)
	fmt.Printf("\n\n")
	fmt.Printf("On try %d out of 6\n", attempt)
	fmt.Printf("Available letters: %v\n", lettersAvailable)
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

func removeStringFromSlice(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func sliceOfString(word string) []string {
	var wordSlice []string
	for _, r := range word {
		// fmt.Println(string(r))
		wordSlice = append(wordSlice, string(r))
	}
	// fmt.Printf("wordSlice: is of type: %v, with value: %v\n", reflect.TypeOf(wordSlice), wordSlice)
	return wordSlice
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func userInput() string {
	fmt.Println("Word guess: ")
	var input string
	fmt.Scanln(&input)
	valid, message := validateWord(input)
	if !valid {
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

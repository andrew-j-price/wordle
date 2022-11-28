package main

import (
	"fmt"
	"math/rand"
	"reflect"
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
	fmt.Printf("guessSlice: is of type: %v, with value: %v\n", reflect.TypeOf(guessSlice), guessSlice)
	fmt.Printf("actualSlice: is of type: %v, with value: %v\n", reflect.TypeOf(actualSlice), actualSlice)
	/*
		for acutal_i, acutal_v := range actualSlice {
			for guess_i, guess_v := range guessSlice {


			}
		}
	*/
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

func populateAvailableLetters() {
	// there probably is a more efficient way
	lettersAvailable = append(lettersAvailable, "a")
	lettersAvailable = append(lettersAvailable, "b")
	lettersAvailable = append(lettersAvailable, "c")
	lettersAvailable = append(lettersAvailable, "d")
	lettersAvailable = append(lettersAvailable, "e")
	lettersAvailable = append(lettersAvailable, "f")
	lettersAvailable = append(lettersAvailable, "g")
	lettersAvailable = append(lettersAvailable, "h")
	lettersAvailable = append(lettersAvailable, "i")
	lettersAvailable = append(lettersAvailable, "j")
	lettersAvailable = append(lettersAvailable, "k")
	lettersAvailable = append(lettersAvailable, "l")
	lettersAvailable = append(lettersAvailable, "m")
	lettersAvailable = append(lettersAvailable, "n")
	lettersAvailable = append(lettersAvailable, "o")
	lettersAvailable = append(lettersAvailable, "p")
	lettersAvailable = append(lettersAvailable, "q")
	lettersAvailable = append(lettersAvailable, "r")
	lettersAvailable = append(lettersAvailable, "s")
	lettersAvailable = append(lettersAvailable, "t")
	lettersAvailable = append(lettersAvailable, "u")
	lettersAvailable = append(lettersAvailable, "v")
	lettersAvailable = append(lettersAvailable, "w")
	lettersAvailable = append(lettersAvailable, "x")
	lettersAvailable = append(lettersAvailable, "y")
	lettersAvailable = append(lettersAvailable, "z")
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

func sliceOfString(word string) []string {
	var wordSlice []string
	for _, r := range word {
		fmt.Println(string(r))
		wordSlice = append(wordSlice, string(r))
	}
	fmt.Printf("wordSlice: is of type: %v, with value: %v\n", reflect.TypeOf(wordSlice), wordSlice)
	return wordSlice
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

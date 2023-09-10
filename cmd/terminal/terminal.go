/*
Package: wordle

An implementation of the word guess game for your Terminal.
*/
package main

import (
	"flag"

	"github.com/andrew-j-price/journey/pkg/logger"
	"github.com/andrew-j-price/wordle/pkg/game"
)

func init() {
	logger.Logger()
}

func main() {
	debugFlag := flag.Bool("debug", false, "include debug output")
	flag.Parse()

	if *debugFlag {
		logger.Debug.Println("Debug mode is set to:", *debugFlag)
	}

	game.GameHandler(*debugFlag)
}

/*
Package: wordle

An implementation of the word guess game for your Terminal or iOS.
*/
package main

import (
	"flag"
	"math/rand"
	"time"
)

var debugFlow bool

func init() {
	simpleLogger()
	rand.Seed(time.Now().Unix())
}

func main() {
	enableDebug := flag.Bool("debug", false, "include debug output")
	flag.Parse()

	debugFlow = *enableDebug
	if debugFlow {
		LoggerInfo.Println("Debug mode is set to:", debugFlow)
	}

	gameHandler()
}

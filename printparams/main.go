package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Write a program that prints the arguments receivedin the command line

func main() {
	arguments := os.Args

	// Function body is called upon 'arguments' above, has to be in main file
	// Convention go run 'x.go' argument , new argument indexed after any space
	// Declare argument in body
	// Loop/iterate to convert string to rune to ascertain values

	for i := 1; i < len(arguments); i++ { // len(args) gives number of arguments
		for _, elements := range arguments[i] {
			z01.PrintRune(elements)
		}
		z01.PrintRune('\n')
	}
}

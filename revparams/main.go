package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args

	for i := len(argument) - 1; i > 0; i-- {
		for _, a := range argument[i] {
			z01.PrintRune(a)
		}
		z01.PrintRune('\n')
	}
}

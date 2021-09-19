package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args

	for i := 1; i < len(argument); i++ {
		for _, a := range argument[i] {
			z01.PrintRune(a)
		}
		z01.PrintRune('\n')
	}
}

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args // creates a variable to hold the value of os.args

	/* slices the argument by first selecting the first item in the array only
	which in this case is the name of program represented as ./main. In order to
	print just the main bit you select only the characters from the 2nd index
	onwards using the 2nd [].*/

	a := argument[0][2:]

	for _, w := range a { // loops through the a variable and prints it as a rune.
		z01.PrintRune(w)
	}
	z01.PrintRune('\n') // prints a new line.
}

package piscine

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	var TRune string = "T"
	var FRune string = "F"
	if nb >= 0 {
		z01.PrintRune(rune(FRune[0]))
	} else {
		z01.PrintRune(rune(TRune[0]))
	}
	z01.PrintRune('\n')
}

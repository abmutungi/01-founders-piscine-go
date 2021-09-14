package piscine

import "github.com/01-edu/z01"

func LastRune(s string) rune {
	var arrayCount rune

	for _, i := range s {
		arrayCount = i
	}
	return arrayCount
}

func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}

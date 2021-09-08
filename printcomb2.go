package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for a := '0'; a <= '8'; a++ {
		for b := '0'; b <= '8'; b++ {
			if a < b {
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(32)
				z01.PrintRune(a)
				for c := '0'; c <= '8'; c++ {
					z01.PrintRune(c)
				}
				if a == '9' && b == '9' {
					z01.PrintRune(10)
				} else {
					z01.PrintRune(32)
				}

			}
		}
	}
}

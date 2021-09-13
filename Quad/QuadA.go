package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) { // 1st line
	if x >= 1 && y >= 1 {
		z01.PrintRune('o')                             // top left corner / starts on first column
		for xcounter := 2; xcounter <= x; xcounter++ { // loop (starts on the 2nd column and ends equal to or lessthan x)
			if xcounter == x { // conditions for the loop (when xcounter is equal to x it will print 'o', if not then print '-')
				z01.PrintRune('o') // top right corner
			} else {
				z01.PrintRune('-') // in the middle
			}
		}
		z01.PrintRune('\n') // new line
	} // 2nd line or middle section
	if y > 2 { // condition for the second IF (when y is greater than 2)
		for ycounter := 2; ycounter < y; ycounter++ { // loop (starts on the 2nd column and ends before y)
			if ycounter <= y { // conditions for the loop (when ycounter is lessthen or equal to y it will print '|')
				z01.PrintRune('|') // left side
			}
			for xcounter := 2; xcounter <= x; xcounter++ { // loop (starts on the 2nd column and ends equal to or lessthan x)
				if xcounter < x { // condition for the loop (when xcounter is lessthan x it will print ' ' Spaces inbetween)
					z01.PrintRune(' ') // the space inbetween
				}
				if xcounter == x { // condition for the loop (when xcounter is equal x it will print '|')
					z01.PrintRune('|') // right side
				}
			}
			z01.PrintRune('\n') // new line
		}
	} // Last line
	if y > 1 { // when y is greater than 1 print 'o'
		z01.PrintRune('o')                             // bottom left corner
		for xcounter := 2; xcounter <= x; xcounter++ { // loop (starts on the 2nd column and ends equal to or lessthan x)
			if xcounter == x { // conditions for the loop (when xcounter is equal to x it will print 'o', if not then print '-')
				z01.PrintRune('o') // bottom right corner
			} else {
				z01.PrintRune('-') // in the middle
			}
		}
		z01.PrintRune('\n') // new line
	}
}

func main() {
	piscine.QuadA(5, 3)
}

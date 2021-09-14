package piscine

func NRune(s string, n int) rune {
	slice := []rune(s) // create a variable called slice

	if len(slice) >= n && n > 0 {
		return slice[n-1]
	}
	return 0
}

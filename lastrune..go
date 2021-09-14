package piscine

func LastRune(s string) rune {
	runeArray := []rune(s)
	arrayCount := 0

	for i := range s {
		arrayCount = i + 1
	}
	return rune(arrayCount - 1)
}

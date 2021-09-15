package piscine

func IsNumeric(s string) bool {
	digit := []rune(s)

	for _, i := range digit {
		if i >= 0 && i <= 47 || i >= 58 {
			return false
		}
	}
	return true
}

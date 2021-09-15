package piscine

func IsAlpha(s string) bool {
	rstr := []rune(s)

	for _, i := range rstr {
		if i >= 0 && i <= 47 || i >= 58 && i <= 64 || i >= 91 && i <= 96 || i >= 123 {
			return false
		}
	}
	return true
}

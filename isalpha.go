package piscine

func IsAlpha(s string) bool {
	rstr := []rune(s)

	for _, i := range rstr {
		if i >= 0 && i <= 32 {
			return false
		}
	}
	return true
}

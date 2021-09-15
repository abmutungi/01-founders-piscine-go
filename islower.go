package piscine

func IsLower(s string) bool {
	rstr := []rune(s)

	for _, i := range rstr {
		if i < 'a' || i > 'z' {
			return false
		}
	}
	return true
}

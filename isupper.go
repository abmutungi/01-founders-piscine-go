package piscine

func IsUpper(s string) bool {
	rstr := []rune(s)

	for _, i := range rstr {
		if i < 'A' || i > 'Z' {
			return false
		}
	}
	return true
}

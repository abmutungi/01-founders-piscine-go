package piscine

func IsUpper(s string) bool {
	rstr := []rune(s)

	for _, i := range rstr {
		if (i >= 'A') && (i <= 'Z') {
			return true
		}
		else if (i >= 'a') && (i <= 'z') {
			return false
		}
}
}
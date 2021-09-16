package piscine

func BasicAtoi(s string) int {
	str := []rune(s)
	var trim int = 0

	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			trim = trim * 10             // trim *=
			trim = trim + int(str[i]-48) // trim += (-48 converts string ascii value to correct number)
		}
	}
	return trim
}

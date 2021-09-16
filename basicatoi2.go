package piscine

func BasicAtoi2(s string) int {
	str := []rune(s)
	var trim int = 0

	for i := 0; i < len(s); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			trim = trim * 10
			trim = trim + int(str[i]-48)
		}
		if (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= 'A' && str[i] <= 'Z') {
			return 0
		}
		if str[i] == ' ' {
			return 0
		}
	}
	return trim
}

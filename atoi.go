package piscine

func Atoi(s string) int {
	str := []rune(s)
	neg := str[0] == '-'
	if neg || str[0] == '+' {
		str = str[1:]
	}
	trim := 0
	for {
		if str[0] >= '0' && str[0] <= '9' {
			trim *= 10
			trim += int(str[0] - 48)
		} else {
			return 0
		}
		if len(str) > 1 {
			str = str[1:]
		} else {
			break
		}
	}
	if neg {
		trim *= -1
	}
	return trim
}

package piscine

func Rot14(s string) string {
	char := []rune(s)

	var result string

	for i := 0; i < len(char); i++ {
		if char[i] >= 'a' && char[i] <= 'z' {
			if char[i] >= 'm' {
				char[i] = char[i] - 12
			} else {
				char[i] = char[i] + 14
			}
		} else if char[i] >= 'A' && char[i] <= 'Z' {
			if char[i] >= 'M' {
				char[i] = char[i] - 12
			} else {
				char[i] = char[i] + 14
			}
		}
		result = result + string(char[i])

	}
	return result
}

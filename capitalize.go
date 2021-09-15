package piscine

func Capitalize(s string) string {
	a := []rune(s) 
	var b string //create extra string

	for i := 0; i <= len(s)-1; i++ {
		if a[i] >= 65 && a[i] <= 90 {
			a[i] = a[i] + 32
		}
		b += string(a[i])
	}
	char := []rune(b)
	c := ""
	for j := 0; j <= len(b)-1; j++ {
		if j == 0 {
			if char[j] >= 97 && char[j] <= 122 {
				char[j] = char[j] - 32
			}
		} else if char[j-1] >= 97 && char[j-1] <= 122 || char[j-1] >= 65 && char[j-1] <= 90 || char[j-1] >= 48 && char[j-1] <= 57 {
		} else if char[j] >= 97 && char[j] <= 122 || char[j] >= 65 && char[j] <= 90 {
			char[j] = char[j] - 32
		}

		c += string(char[j])
	}
	return c
}

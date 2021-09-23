package piscine

func Capitalize(s string) string {
	a := []rune(s) // slicing the string s

	for i := 0; i < len(s); i++ { // can also be i <= len(s)-1
		if a[i] >= 'A' && a[i] <= 'Z' { // if capitalised then...
			a[i] = a[i] + 32 //...make lower case
		}
	}
	char := []rune(string(a)) // slicing the string created above
	for j := 0; j < len(s); j++ {
		if j == 0 { // capitalising the first character if its a lowercase letter
			if char[j] >= 'a' && char[j] <= 'z' {
				char[j] = char[j] - 32
			} // all the conditions for capitalising if its preceded by a character
		} else if char[j-1] >= 'a' && char[j-1] <= 'z' || char[j-1] >= 'A' && char[j-1] <= 'Z' || char[j-1] >= '0' && char[j-1] <= '9' {
		} else {
			char[j] = char[j] - 32
		}
	}
	return string(char)
}

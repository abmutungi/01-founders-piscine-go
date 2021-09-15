package piscine

func Capitalize(s string) string {
	a := []rune(s) // slicing the string s

	for i := 0; i < len(s); i++ { // can also be i <= len(s)-1
		if a[i] >= 65 && a[i] <= 90 { // if a capitalised then...
			a[i] = a[i] + 32 //...make lower case
		}
	}
	char := []rune(string(a)) // slicing the string created above
	for j := 0; j < len(s); j++ {
		if j == 0 { // capitalising the first character if its a lowercase letter
			if char[j] >= 97 && char[j] <= 122 {
				char[j] = char[j] - 32
			}
		} else if char[j-1] >= 97 && char[j-1] <= 122 || char[j-1] >= 65 && char[j-1] <= 90 || char[j-1] >= 48 && char[j-1] <= 57 {
		} else if char[j] >= 97 && char[j] <= 122 || char[j] >= 65 && char[j] <= 90 {
			char[j] = char[j] - 32
		}
	}
	return string(char)
}

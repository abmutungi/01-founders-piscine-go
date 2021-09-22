func Rot14(s string) string {

	char := []rune(s)

	for _, char:= range s {
		if char >= 65 && char <= 76 || char >= 97 && char <= 108 {
		result = append(result, char+14)
	}
		if char >= 77 && char <= 90 || char >= 109 && char <= 122 {
		result = append(result, char-12)

		} else { return char

}
return result
	}
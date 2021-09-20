package piscine

func SplitWhiteSpaces(str string) []string {
	runes := []rune(str) // break string into runes

	var answer []string // creat empty string array

	result := "" // intialise result with empty string

	for i := 0; i < len(runes); i++ {
		if runes[i] == ' ' || runes[i] == '\n' || runes[i] == '\t' {
			if result != "" { // result updates every interation
				answer = append(answer, result) // append previous result
				result = ""                     // update result to be ""
			}
		} else {
			result = result + string(runes[i])
		}
		if i == len(runes)-1 {
			answer = append(answer, result) // append last result
		}
	}
	return answer
}

package piscine

import "strings"

func SplitWhiteSpaces(str string) []string {
	split := strings.Fields(str)
	return split
}

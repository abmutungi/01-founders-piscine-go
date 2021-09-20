package piscine

import "strings"

func SplitWhiteSpaces(str string) []string {
	space := strings.Fields(str)
	return space
}

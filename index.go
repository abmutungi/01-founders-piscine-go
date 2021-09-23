package piscine

func Index(s string, toFind string) int {
	/*	place := []rune(s)
		find := len([]rune(toFind))

		if len(place) < find {
			return -1
		}*/
	for i := range s {
		if len(s[i:]) < len(toFind) {
			return -1
		} else {
			if string(s[i:i+len(toFind)]) == toFind {
				return i
			}
		}
	}
	return -1
}

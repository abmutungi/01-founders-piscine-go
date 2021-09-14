package piscine

func Index(s string, toFind string) int {
	place := []rune(s)
	find := len([]rune(toFind))

	if len(place) < find {
		return -1
	}
	for i := 0; i < len(place); i++ {
		if len(place[i:]) < len([]rune(toFind)) {
			return -1
		} else {
			if s[i:i+find] == toFind {
				return i
			}
		}
	}
	return -1
}

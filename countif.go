package piscine

func CountIf(f func(string) bool, tab []string) int {
	var nb int = 0

	for _, howManyNum := range tab {
		if f(howManyNum) {
			nb++
		}
	}
	return nb
}

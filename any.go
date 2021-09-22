package piscine

func Any(f func(string) bool, a []string) bool {
	for _, isNum := range a {
		if f(isNum) {
			return true
		}
	}
	return false
}

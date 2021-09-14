package piscine

func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a > b { // words are the same except for last letters
		return 1
	} // everythging else

	return -1
}

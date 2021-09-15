package piscine

func IsPrintable(s string) bool {
	for _, i := range s {
		if i >= 0 && i <= 31 || i >= 127 {
			return false
		}
	}
	return true
}

package piscine

func ToUpper(s string) string {
	a := []rune(s)

	for i := 0; i < len(s); i++ {
		if a[i] >= 'a' && a[i] <= 'z' {
			a[i] = a[i] - 32
		}
	}
	return string(a)
}

package piscine

func Swap2(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}

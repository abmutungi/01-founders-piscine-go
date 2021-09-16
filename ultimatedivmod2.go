package piscine

func UltimateDivMod2(a *int, b *int) {
	c := *a
	d := *b
	*a = c / d
	*b = c % d
}

package piscine

func DivMod2(a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b
}

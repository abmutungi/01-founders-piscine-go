package piscine

func StrRev2(s string) string {
	srune := []rune(s)
	var arr []rune

	for i := len(s) - 1; i >= 0; i-- {
		arr = append(arr, srune[i])
	}
	return string(arr)
}

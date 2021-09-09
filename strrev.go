package piscine

func StrRev(s string) string {
	conv := []rune(s)

	var arr []rune

	for i := len(conv) - 1; i > -1; i-- {
		arr = append(arr, conv[i])
	}
	return string(arr)
}

package piscine

func ConcatParams(args []string) string {
	str := ""

	for i, result := range args {
		str += string(result)
		if i != len(args)-1 {
			str += "\n"
		}
	}
	return str
}

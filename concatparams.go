package piscine

func ConcatParams(args []string) string {
	var result string
	for index, a := range args {
		if index == 0 {
			result = a
		}
		result += "\n" + a
	}
	return result
}

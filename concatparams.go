package piscine

func ConcatParams(args []string) string {
	var argument string = args[0] // 0 to print the 1st string
	for _, i := range args[1:] {  // range loop args[1:] is set to print the following arguments
		argument += "\n" + i // (+=) is add and assign- add var (a) with new lines and the other arguments
	} // "\n" has to go first before adding (s) for it to print each argument on a separate line
	return argument
}
package piscine

func Atoi(s string) int {
	num := 0    // initalise num as 0
	result := 1 // initialise result as 1

	for i, val := range s { // uses for range loop to iterate over s
		digit := int(val) - 48
		if digit <= 9 && digit >= 0 {
			num = num*10 + digit
		} else if digit == -3 && i == 0 {
			result = -1
		} else if digit == -5 && i == 0 {
			result = 1
		} else {
			return 0
		}
	}
	num *= result
	return num
}

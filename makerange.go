package piscine

func MakeRange(min, max int) []int {
	var result []int = make([]int, max-min)

	for i := 0; i < (max - min); i++ {
		result[i] = i + min
		if max <= min {
			return nil
		}

	}
	return result
}

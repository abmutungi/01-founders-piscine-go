package piscine

func MakeRange(min, max int) []int {
	var result []int = make([]int, max-min)

	if max <= min {
		return nil
	}
	for i := 0; i < (max - min); i++ {
		result[i] = i + min
	}
	return result
}

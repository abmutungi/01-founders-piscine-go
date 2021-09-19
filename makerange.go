package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		return nil
	}
	result := make([]int, max-min)

	for i := range result {
		result[i] = i + min
	}
	return result
}

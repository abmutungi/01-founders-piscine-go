package piscine

func IterativePower(nb int, power int) int {
	result := 1
	if power < 0 {
		return 0
	}

	if power == 1 {
		return nb
	}

	if power > 1 && nb >= 1 {
		for i := 1; i <= power; i++ {
			result = result * nb
		}
	}
	return result
}

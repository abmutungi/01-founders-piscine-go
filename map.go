package piscine

func Map(f func(int) bool, a []int) []bool {
	var result []bool

	for _, nbr := range a {
		result = append(result, f(nbr))
	}
}

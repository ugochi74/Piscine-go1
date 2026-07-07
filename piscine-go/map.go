package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, r := range a {
		result[i] = f(r)
	}
	return result
}

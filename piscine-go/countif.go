package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, d := range tab {
		if f(d) {
			count++
		}
	}
	return count
}

package piscine

func Any(f func(string) bool, a []string) bool {
	for _, r := range a {
		if f(r) {
			return true
		}
	}
	return false
}

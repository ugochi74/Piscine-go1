package piscine

func IsUpper(s string) bool {
	r := []rune(s)
	for i := range r {
		if r[i] < 60 || r[i] > 90 {
			return false
		}
	}
	return true
}

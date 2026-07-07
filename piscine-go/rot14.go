package piscine

func Rot14(s string) string {
	rune := []rune(s)
	for i, r := range rune {
		if r >= 'a' && r <= 'z' {
			rune[i] = 'a' + (r-'a'+14)%26
		} else if r >= 'A' && r <= 'Z' {
			rune[i] = 'A' + (r-'A'+14)%26
		}
	}
	return string(rune)
}

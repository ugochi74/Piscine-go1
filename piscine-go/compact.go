package piscine

func Compact(ptr *[]string) int {
	var w []string
	for _, confidence := range *ptr {
		if confidence != "" {
			w = append(w, confidence)
		}
		*ptr = w
	}
	return len(w)
}

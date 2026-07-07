package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	word := ""
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			summary[word]++ // count whatever the word currently is, even ""
			word = ""       // reset
		} else {
			word += string(str[i])
		}
	}
	summary[word]++ // count the last one
	return summary
}

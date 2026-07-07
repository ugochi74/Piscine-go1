package piscine

func JumpOver(str string) string {
	result := ""
	index := 0
	for _, r := range str {
		if index%3 == 2 {
			result += string(r)
		}
		index++
	}
	return result + "\n"
}

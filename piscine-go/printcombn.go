package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	digits := []int{}
	for i := 0; i < n; i++ {
		digits = append(digits, i)
	}
	for {
		for i := 0; i < n; i++ {
			z01.PrintRune(rune(digits[i] + 48))
		}
		pos := n - 1
		for pos >= 0 && digits[pos] == 9-(n-1-pos) {
			pos--
		}
		if pos < 0 {
			break
		}
		z01.PrintRune(',')
		z01.PrintRune(' ')
		digits[pos]++
		for i := pos + 1; i < n; i++ {
			digits[i] = digits[i-1] + 1
		}
	}
	z01.PrintRune('\n')
}

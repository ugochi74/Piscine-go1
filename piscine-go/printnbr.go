package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	sliceNB := []int{}
	i := 0
	isMinInt := false
	if n < 0 {
		z01.PrintRune('-')
		if n == -9223372036854775808 {
			n = n + 1
			isMinInt = true
		}
		n = -n
	} else if n == 0 {
		z01.PrintRune('0')
	}
	for n > 0 {
		sliceNB = append(sliceNB, n%10)
		n = n / 10
		i++
	}
	for j := len(sliceNB) - 1; j >= 0; j-- {
		if isMinInt && j == 0 {
			z01.PrintRune('8')
		} else {
			z01.PrintRune(rune(sliceNB[j] + 48))
		}
	}
}

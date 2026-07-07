package piscine

import "fmt"

func QuadA(x, y int) {

	if x <= 0 || y <= 0 {
		return
	}

	for r := 1; r <= y; r++ {
		for c := 1; c <= x; c++ {
			if (r == 1 || r == y) && (c == 1 || c == x) {
				fmt.Print("o")

			} else if r == 1 || r == y {
				fmt.Print("-")

			} else if c == 1 || c == x {
				fmt.Print("|")

			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

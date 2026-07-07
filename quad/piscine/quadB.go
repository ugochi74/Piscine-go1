package piscine

import "fmt"

func QuadB(x, y int) {
	if x < 0 || y < 0 {
		return
	}

	for r := 1; r <= y; r++ {
		for c := 1; c <= x; c++ {
			if r == 1 && c == 1 {
				fmt.Print("/")

			} else if r == 1 && c == x {
				fmt.Print("\\")
			} else if r == y && c == 1 {
				fmt.Print("\\")

			} else if r == y && c == x {
				fmt.Print("/")

			} else if r == 1 || r == y || c == 1 || c == x {
				fmt.Print("*")

			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

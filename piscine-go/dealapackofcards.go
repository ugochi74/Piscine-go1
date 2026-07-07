package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	for player := 1; player <= 4; player++ {
		fmt.Printf("Player %d: ", player)

		start := (player - 1) * 3
		end := start + 3

		for i := start; i < end; i++ {
			if i == end-1 {
				fmt.Printf("%d", deck[i])
			} else {
				fmt.Printf("%d, ", deck[i])
			}
		}

		fmt.Printf("\n")
	}
}

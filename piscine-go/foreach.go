package piscine

func ForEach(f func(int), a []int) {
	for _, r := range a {
		f(r)
	}
}

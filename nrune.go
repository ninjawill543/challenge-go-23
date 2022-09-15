package piscine

func NRune(s string, n int) rune {
	var count []int
	for _, c := range s {
		count = append(count, int(c))
	}
	if n < 1 {
		return 0
	} else {
		return rune(count[(n - 1)])
	}
}

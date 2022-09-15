package piscine

func NRune(s string, n int) rune {
	var count []int
	for _, c := range s {
		count = append(count, int(c))
	}
	if n < 1 || n > (len(count)) {
		return 0
	} else {
		return rune(count[(n - 1)])
	}
}

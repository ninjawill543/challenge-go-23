package piscine

func LastRune(s string) rune {
	var count []int
	for _, c := range s {
		count = append(count, int(c))
	}
	return rune(count[(len(count) - 1)])
}

package piscine

func FirstRune(s string) rune {
	var count []int
	for _, c := range s {
		count = append(count, int(c))
	}
	return rune(count[0])
}

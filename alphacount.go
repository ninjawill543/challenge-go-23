package piscine

func AlphaCount(s string) int {
	var count []int
	for _, c := range s {
		if (c >= 65 && c <= 90) || (c >= 97 && c <= 122) {
			count = append(count, int(c))
		}
	}
	return (len(count))
}

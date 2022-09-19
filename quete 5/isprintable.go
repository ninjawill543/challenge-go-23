package piscine

func IsPrintable(s string) bool {
	count := 0
	for _, c := range s {
		if c >= 32 && c <= 127 {
			count++
		}
		if count == len(s) {
			return true
		}
	}
	return false
}

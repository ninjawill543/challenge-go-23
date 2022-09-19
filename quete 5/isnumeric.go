package piscine

func IsNumeric(s string) bool {
	count := 0
	for _, c := range s {
		if c >= 48 && c <= 57 {
			count++
		}
		if count == len(s) {
			return true
		}
	}
	return false
}

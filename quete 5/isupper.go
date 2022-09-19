package piscine

func IsUpper(s string) bool {
	count := 0
	for _, c := range s {
		if c >= 65 && c <= 90 {
			count++
		}
		if count == len(s) {
			return true
		}
	}
	return false
}

package piscine

func IsLower(s string) bool {
	count := 0
	for _, c := range s {
		if c >= 97 && c <= 122 {
			count++
		}
		if count == len(s) {
			return true
		}
	}
	return false
}

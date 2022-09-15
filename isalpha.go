package piscine

func IsAlpha(s string) bool {
	count := 0
	for _, c := range s {
		if c >= 97 && c <= 122 || c >= 48 && c <= 57 || c >= 65 && c <= 90 {
			count++
		}
		if count == len(s) {
			return true
		}
	}
	return false
}

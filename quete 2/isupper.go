package piscine

func IsUpper(s string) bool {
	//var count []int
	for _, c := range s {
		if c >= 65 && c <= 90 {
			return true
		}
	}
	return false
}

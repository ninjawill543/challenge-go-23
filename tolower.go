package piscine

func ToLower(s string) string {
	var count string
	for _, c := range s {
		if c >= 32 && c <= 64 || c >= 91 && c <= 127 {
			count += string(c)
		}
		if c >= 65 && c <= 90 {
			numb := c + 32
			count += (string(numb))
		}

	}
	return count
}

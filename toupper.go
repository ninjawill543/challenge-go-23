package piscine

func ToUpper(s string) string {
	var count string
	for _, c := range s {
		if c >= 32 && c <= 96 || c >= 123 && c <= 127 {
			count += string(c)
		}
		if c >= 97 && c <= 122 {
			numb := c - 32
			count += (string(numb))
		}

	}
	return count
}

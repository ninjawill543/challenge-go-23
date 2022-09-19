package piscine

func Capitalize(s string) string {
	counter := 0
	const A_low = 
	const A_high = 
	const 
	var count string
	for _, c := range s {
		if counter == 0 {
			if c >= 32 && c <= 96 || c >= 123 && c <= 127 {
				counter++
				count += string(c)
			}
			if c >= 97 && c <= 122 {
				counter++
				numb := c - 32
				count += (string(numb))
			}
		}
		if c >= 32 && c <= 64 || c >= 91 && c <= 127 {
			count += string(c)
			counter++
		}
		if c >= 65 && c <= 90 {
			numb := c + 32
			count += (string(numb))
			counter++
		}
		if c < 48 || c > 57 && c < 65 || c > 90 && c < 97 || c > 122 {
			counter = 0
			count += string(c)
		}

	}
	return count
}

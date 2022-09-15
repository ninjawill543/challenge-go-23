package piscine

func Index(s string, toFind string) int {
	counta := []rune(s)
	countb := []rune(toFind)
	counter := 0
	cb := 0
	for i := 0; i < len(s); i++ {
		if counta[i] != countb[cb] {
			counter++
		} else if counta[i] == countb[cb] {
			cb++
			counter++
		}
	}

	return counter

}

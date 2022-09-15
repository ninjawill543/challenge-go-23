package piscine

func Index(s string, toFind string) int {
	counta := []rune(s)
	countb := []rune(toFind)
	counter := 0
	ca := 0
	cb := 0
	if counta[ca] == countb[cb] {
		ca++
		cb++
		counter++
	} else if counta[ca] != countb[cb] {
		ca++
		counter++
	}
	//exit:
	return counter

}

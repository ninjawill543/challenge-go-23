package piscine

func IterativeFactorial(nb int) int {
	result := 1
	for i := 1; i <= nb; i++ {
		if nb > 0 {
			result = result * i
		} else {
			return 0
		}
	}
	return result
}

package piscine

func IterativeFactorial(nb int) int {
	result := 1
	for i := 1; i <= nb; i++ {
		if nb < 0 || nb > 10 {
			result = 0
		} else {
			result = result * i
		}
	}
	return result
}

package piscine

func IterativeFactorial(nb int) int {
	result := 1
	for i := 1; i <= nb; i++ {
		if nb > 0 && nb < 10 {
			result = result * i
		} else if nb < 0 {
			result = 0
		}
	}
	return result
}

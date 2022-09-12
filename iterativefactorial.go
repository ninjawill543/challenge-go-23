package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb < 0 {
		result = 0
		goto Exit
	} else {
		for i := 1; i <= nb; i++ {
			if nb > 0 && nb < 10 {
				result = result * i
			} else {
				result = 0
			}
		}
	}
Exit:
	return result
}

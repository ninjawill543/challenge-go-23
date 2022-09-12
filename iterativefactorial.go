package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb > 0 && nb < 10 {
		for i := 1; i <= nb; i++ {
			result = result * i
		}
		goto Exit
	} else {
		result = 0
		goto Exit
	}
Exit:
	return result
}

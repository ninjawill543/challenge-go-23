package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb > 0 && nb < 100 {
		for i := 1; i <= nb; i++ {
			result = result * i
		}
		goto Exit
	} else if nb == 0 {
		result = 1
		goto Exit
	} else {
		result = 0
		goto Exit
	}
Exit:
	return result
}

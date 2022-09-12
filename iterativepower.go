package piscine

func IterativePower(nb int, power int) int {
	result := 1
	if nb > 0 && nb < 100 {
		for i := 0; i < power; i++ {
			result = result * nb
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

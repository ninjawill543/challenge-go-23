package piscine

func IterativePower(nb int, power int) int {
	result := 1
	if power > 0 && power < 100 {
		for i := 0; i < power; i++ {
			result = result * nb
		}
		goto Exit
	} else if power == 0 {
		result = 1
		goto Exit
	} else {
		result = 0
		goto Exit
	}
Exit:
	return result
}

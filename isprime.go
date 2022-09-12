package piscine

func IsPrime(nb int) bool {
	for i := 2; i < 100; i++ {
		if nb%i == 0 {
			return false
		} else if nb%i != 0 {
			return true
		}
	}
	return false
}

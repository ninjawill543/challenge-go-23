package piscine

func IsPrime(nb int) bool {
	for i := 1; i < 100; i++ {
		result := nb / i
		if result*i == nb {
			return false
		}
	}
	return true
}

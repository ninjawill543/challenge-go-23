package piscine

func IsPrime(nb int) bool {
	for i := 2; i < 1000000; i++ {
		if nb > 0 {
			if nb%i == 0 {
				return false
			} else if nb%i != 0 {
				return true
			}
		} else {
			return false
		}
	}
	return false
}

package piscine

func IsPrime(nb int) bool {
	for i := 2; i < 1000000; i++ {
		if nb%i == 0 && nb > 0 {
			return false
		}
	}
	if nb > 0 {
		return true
	} else {
		return false
	}
}

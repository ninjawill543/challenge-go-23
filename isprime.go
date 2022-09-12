package piscine

func IsPrime(nb int) bool {
	for i := 2; i < nb; i++ {
		if nb < 0 {
			return false
		} else if nb%i == 0 && nb > 0 {
			return false
		}
	}
	if nb > 0 {
		return true
	}
}

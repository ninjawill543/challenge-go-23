package piscine

func Sqrt(nb int) int {
	resultat := 0
	for i := 1; i < 1000; i++ {
		if nb == i*i {
			resultat = i
		} else {
			return 0
		}
	}

	return resultat
}

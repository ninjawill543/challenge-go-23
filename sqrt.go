package piscine

func Sqrt(nb int) int {
	resultat := 0
	for i := 1; i < 1000; i++ {
		if nb == i*i {
			resultat = i
		} else if nb == 1 {
			return 1
		}
	}

	return resultat
}

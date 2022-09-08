package piscine

func StrLen(s string) int {
	cpt := 0
	liste := []rune{}
	for _, lettre := range s {
		liste = append(liste, lettre)
		cpt++
	}
	return cpt
}

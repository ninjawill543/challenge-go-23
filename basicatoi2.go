package piscine

func BasicAtoi2(s string) int {
	reponse := 0
	n := 1
	list := []rune{}
	for _, verif := range s {
		if verif != '0' && verif != '1' && verif != '2' && verif != '3' && verif != '4' && verif != '5' && verif != '6' && verif != '7' && verif != '8' && verif != '9' {
			return 0
		} else {
			list = append(list, verif)
		}
	}
	for i := len(list) - 1; len(list) != 0; i-- {
		reponse += int(list[i]-48) * n
		list = list[:len(list)-1]
		n *= 10
	}
	return reponse
}

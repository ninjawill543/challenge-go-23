package piscine

func SplitWhiteSpaces(s string) []string {
	list := []string{}
	hold := ""
	for _, i := range s {
		if i != 9 && i != 32 && i != 10 {
			hold = hold + string(i)
		} else {
			list = append(list, hold)
			hold = ""
		}
	}
	list = append(list, hold)
	return list
}

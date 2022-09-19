package piscine

func SplitWhiteSpaces(s string) []string {
	list := []string{}
	holder := ""
	for _, i := range s {
		if i != 9 && i != 32 && i != 10 {
			holder = holder + string(i)
		} else {
			list = append(list, holder)
			holder = ""
		}
	}
	list = append(list, holder)
	return list
}

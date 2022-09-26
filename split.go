package piscine

func Split(s, sep string) []string {
	list := []string{}
	hold := ""
	for _, m := range s {
		if m != 9 && m != 32 && m != 10 {
			hold = hold + string(m)
		} else {
			if hold != "" {
				list = append(list, hold)
				hold = ""
			}
		}
	}
	list = append(list, hold)
	return list
}

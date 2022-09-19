package piscine

func ConcatParams(args []string) string {
	var list string
	counter := 0
	for _, i := range args {
		counter++
		list = list + i
		if counter == len(args) {
			return list
		}
		list = list + "\n"
	}

	return list
}

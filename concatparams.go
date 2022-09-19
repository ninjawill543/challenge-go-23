package piscine

func ConcatParams(args []string) string {
	var list string
	for _, i := range args {
		list = list + i
		list = list + "\n"
	}

	return list
}

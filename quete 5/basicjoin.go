package piscine

func BasicJoin(elems []string) string {
	var word string
	for _, c := range elems {
		word = word + c
	}
	return word
}

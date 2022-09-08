package piscine

func Swap(a *int, b *int) {
	var c int = *b
	*b = *a
	*a = c
}

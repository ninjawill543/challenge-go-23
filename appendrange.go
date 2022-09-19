package piscine

func AppendRange(min, max int) []int {
	list := []int{}
	if min < max {
		for i := min; i < max; i++ {
			list = append(list, i)
		}
	}
	return list
}

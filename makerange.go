package piscine

func MakeRange(min, max int) []int {
	list := []int{}
	if min < max {
		for i := min; i < max; i++ {
			list[i] = i + i
		}
	} else {
		list = nil
	}
	return list
}

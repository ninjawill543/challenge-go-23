package piscine

func MakeRange(min int, max int) []int {
	list := []int(nil)
	if min < max {
		list := make([]int, (max - min))
		for i := min; i < max; i++ {
			list[i-min] = i
		}
	}
	return list
}

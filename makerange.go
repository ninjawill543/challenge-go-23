package piscine

func MakeRange(min, max int) []int {
	list := []int(nil)
	if min < max {
		list := make([]int, (max - min))
		for i := min; i < max; i++ {
			list[i-min] = i + 1
		}
	} else {
		list = nil
	}
	return list
}

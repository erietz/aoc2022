package aoc

func MinMax(array []int) (int, int) {
	min := array[0]
	max := array[0]
	for _, v := range array {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}


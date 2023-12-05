package maths

func Min[T NumberType](nums ...T) (min T) {
	min = nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

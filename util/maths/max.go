package maths

func Max[T NumberType](nums ...T) (max T) {
	max = nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

package maths

// Function to calculate LCM (Least Common Multiple)
func LCM(nums ...int) int {
	lcm := nums[0]
	for i := 1; i < len(nums); i++ {
		lcm = lcm * nums[i] / GCD(lcm, nums[i])
	}
	return lcm
}

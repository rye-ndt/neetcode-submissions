// should i rob house i?
func rob(nums []int) int {	
	if len(nums) == 1 {
		return nums[0]
	}
	
	note := make([]int, len(nums))
	note[0] = nums[0]
	note[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		note[i] = max(
			note[i-1], note[i-2] + nums[i],
		)
	}

	return note[len(nums)-1]
}

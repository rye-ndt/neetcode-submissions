func rob(nums []int) int {	
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	note := make([]int, len(nums))
	note[0] = nums[0]
	note[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		note[i] = max(note[i-2] + nums[i], note[i-1])
	}

	return note[len(nums)-1]
}

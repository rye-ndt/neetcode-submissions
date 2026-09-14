func rob(nums []int) int {	
	if len(nums) < 3 { return max(nums[0], nums[len(nums)-1]) }
	robs := make([]int, len(nums))
	robs[0], robs[1] = nums[0], max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		robs[i] = max(nums[i] + robs[i-2], robs[i-1])
	}

	return robs[len(robs)-1]
}
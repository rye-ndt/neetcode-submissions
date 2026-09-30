func rob(nums []int) int {	
	if len(nums) < 3 { return max(nums[0], nums[len(nums)-1]) }

	walk := make([]int, len(nums))
	walk[0], walk[1] = nums[0], max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		walk[i] = max(nums[i] + walk[i-2], walk[i-1])
	}

	return walk[len(walk)-1]
}
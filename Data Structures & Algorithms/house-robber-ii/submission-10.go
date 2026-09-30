func rob(nums []int) int {
	if len(nums) < 3 { return max(nums[0], nums[len(nums)-1]) }
	return max(helper(nums[1:]), helper(nums[:len(nums)-1]))
}

func helper(nums []int) int {
	walk := make([]int, len(nums))
	walk[0], walk[1] = nums[0], max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		walk[i] = max(walk[i-1], walk[i-2] + nums[i])
	}

	return walk[len(walk)-1]
}
func rob(nums []int) int {	
	switch len(nums) {
		case 1: return nums[0]
		case 2: return max(nums[0], nums[1])
		default: 
			robs := make([]int, len(nums))
			robs[0], robs[1] = nums[0], max(nums[0], nums[1])

			for i := 2; i < len(nums); i++ {
				robs[i] = max(nums[i] + robs[i-2], robs[i-1])
			}

			return robs[len(robs)-1]
	}
}
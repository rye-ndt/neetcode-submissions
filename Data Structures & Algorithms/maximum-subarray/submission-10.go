func maxSubArray(nums []int) int {
	cur, best := nums[0], nums[0]

	for _, n := range nums[1:] {
		cur = max(n, cur + n)
		best = max(best, cur)
	}

	return best 
}
// at position i:
// should i extend the previous array, or start fresh?
// -> keep, or reset

func maxSubArray(nums []int) int {
	cur := nums[0]
	best := nums[0]

	for i := 1; i < len(nums); i++ {
		cur = max(nums[i], nums[i] + cur)
		best = max(best, cur)
	}

	return best
}
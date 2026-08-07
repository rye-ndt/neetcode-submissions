// should i rob house i?
func rob(nums []int) int {	
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	
	note := make([]int, n)
	note[0] = nums[0]
	note[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		note[i] = max(
			nums[i] + note[i-2],
			note[i-1],
		)
	}

	return note[n-1]
}

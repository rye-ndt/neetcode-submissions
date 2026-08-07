func maxRound(nums []int) int {
	prev2 := nums[0]
	prev1 := max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		cur := max(
			prev1,
			prev2 + nums[i],
		)

		prev2 = prev1
		prev1 = cur
	}

	return prev1
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	return max(
		maxRound(nums[1:]), 
		maxRound(nums[:len(nums)-1]),
	)
}

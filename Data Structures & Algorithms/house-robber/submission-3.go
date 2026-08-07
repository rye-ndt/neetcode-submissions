func rob(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}

	if len(nums) < 3 {
		return max(nums[0], nums[1])
	}

    note := make([]int, len(nums))

	// should i rob house i, or skip it?
	// if i do, i must skip i-1, and take i-2 only 
	// else, i have nothing here
	note[0] =  nums[0]
	note[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		// should i rob at i?
		note[i] = max(nums[i] + note[i-2], note[i-1])
	}

	return note[len(nums)-1]
}

func eraseOverlapIntervals(nums [][]int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})

	prevEnd, counter := nums[0][1], 0

	for i := 1; i < len(nums); i++ {
		if nums[i][0] >= prevEnd {
			prevEnd = nums[i][1]
			continue 
		} 
		
		// end early -> more room -> less overlap
		prevEnd = min(prevEnd, nums[i][1]) 
		counter++
	}

	return counter
}
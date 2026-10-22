func eraseOverlapIntervals(nums [][]int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})

	fmt.Println(nums)

	prevEnd := nums[0][1]
	counter := 0

	for i := 1; i < len(nums); i++ {
		cur := nums[i]

		if cur[0] >= prevEnd {
			prevEnd = cur[1]
			continue 
		}
		
		prevEnd = min(prevEnd, cur[1])
		counter++
	}

	return counter
}
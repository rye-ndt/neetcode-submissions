func eraseOverlapIntervals(nums [][]int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][1] < nums[j][1]
	})

	count, end := 0, math.MinInt
	for _, iv := range nums {
		if iv[0] >= end {
			end = iv[1]
		} else {
			count++
		}
	}
	return count
}
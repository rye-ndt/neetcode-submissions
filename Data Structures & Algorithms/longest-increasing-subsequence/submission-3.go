func lengthOfLIS(nums []int) int {
	n := len(nums)
	note := make([]int, n)
	best := 0

	for i := 0; i < n; i++ {
		note[i] = 1

		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && note[j]+1 > note[i] {
				note[i] = note[j]+1
			}
		}

		if note[i] > best {
			best = note[i]
		}
	}

	return best
}

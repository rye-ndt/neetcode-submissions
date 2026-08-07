func lengthOfLIS(nums []int) int {
    n := len(nums)
	note := make([]int, n)
	best := 0

	for i := range nums {
		note[i] = 1

		for j := 0; j < i; j++ {
			better := note[j] + 1 > note[i]
			smaller := nums[j] < nums[i]

			if smaller && better {
				note[i] = note[j] + 1
			}
		}

		if note[i] > best {
			best = note[i]
		}
	}

	return best
}

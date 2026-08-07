func lengthOfLIS(nums []int) int {
    note := make([]int, len(nums))
	best := 0

	for i := 0; i < len(nums); i++ {
		note[i] = 1

		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] {
				continue
			}

			if note[j]+1 > note[i] {
				note[i] = note[j]+1
			}
		}

		if note[i] > best {
			best = note[i]
		}
	}

	return best
}

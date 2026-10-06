import "slices"

func lengthOfLIS(nums []int) int {
	walk := slices.Repeat([]int{1}, len(nums))
	best := 0

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] { continue }

			walk[i] = max(walk[i], walk[j]+1)
		}

		best = max(best, walk[i])
	}

	return best
}

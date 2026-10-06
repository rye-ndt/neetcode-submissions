import "slices"

func lengthOfLIS(nums []int) int {
	walk := slices.Repeat([]int{1}, len(nums))

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] { continue }

			walk[i] = max(walk[i], walk[j]+1)
		}
	}

	sort.Ints(walk)

	return walk[len(walk)-1]
}

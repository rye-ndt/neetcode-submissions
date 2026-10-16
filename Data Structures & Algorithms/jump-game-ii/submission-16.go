func jump(nums []int) int {
	i, counter, n := 0, 0, len(nums)

	for i < n-1 {
		maxIndex := i+1

		for j := i+1; j <= i + nums[i] && j < n; j++ {
			if min(j + nums[j], n-1) >= min(maxIndex + nums[maxIndex], n-1) {
				maxIndex = j
			}
		}

		i = maxIndex
		counter++
	}

	return counter
}
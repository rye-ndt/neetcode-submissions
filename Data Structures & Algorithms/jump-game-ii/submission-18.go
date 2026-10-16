func jump(nums []int) int {
	counter, n := 0, len(nums)

	for i := 0; i < n-1; i++ {
		maxIndex := i+1

		for j := i+1; j <= i + nums[i] && j < n; j++ {
			if min(j + nums[j], n-1) >= min(maxIndex + nums[maxIndex], n-1) {
				maxIndex = j
			}
		}

		i = maxIndex-1
		counter++
	}

	return counter
}
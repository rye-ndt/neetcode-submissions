func jump(nums []int) int {
	i, counter := 0, 0

	for i < len(nums)-1 {
		cur := nums[i]
		maxIndex := i+1

		for j := i+1; j <= i + cur && j < len(nums); j++ {
			if min(j + nums[j], len(nums)-1) >= min(maxIndex + nums[maxIndex], len(nums)-1) {
				maxIndex = j
			}
		}

		i = maxIndex
		counter++
	}

	return counter
}
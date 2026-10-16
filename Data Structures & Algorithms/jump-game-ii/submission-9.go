func jump(nums []int) int {
	i, counter := 0, 0

	for i < len(nums)-1 {
		cur := nums[i]
		maxIndex := 0
		tempMax := 0

		for j := i+1; j <= i + cur && j < len(nums); j++ {
			temp := min(j + nums[j], len(nums)-1)

			if temp >= tempMax {
				tempMax = temp
				maxIndex = j
			}
		}

		i = maxIndex
		counter++
	}

	return counter
}
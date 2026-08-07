func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

    counter := 0
	i := 0

	for i < len(nums) {
		if i + nums[i] >= len(nums) - 1 {
			return counter+1
		}

		localMax := nums[i]
		maxJ := i

		for j := i+1; j <= i + nums[i]; j++ {
			if nums[i] + nums[j] + j >= localMax {
				localMax = nums[i] + nums[j] + j 
				maxJ = j
			}
		}

		counter++
		i = maxJ 	
	}

	return counter
}

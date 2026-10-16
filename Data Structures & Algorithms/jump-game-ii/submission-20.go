func jump(nums []int) int {
	counter, n := 0, len(nums)

	for i := 0; i < n-1; i++ {
		chose := i+1

		for j := chose; j <= i + nums[i] && j < n; j++ {
			if min(j + nums[j], n-1) >= min(chose + nums[chose], n-1) {
				chose = j
			}
		}

		i = chose-1
		counter++
	}

	return counter
}
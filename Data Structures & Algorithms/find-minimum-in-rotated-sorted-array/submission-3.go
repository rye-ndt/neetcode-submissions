func findMin(nums []int) int {
	l, r := 0, len(nums)-1

	result := nums[0]

	for l <= r {
		m := (r + l) / 2

		result = min(result, nums[m])

		if nums[m] > nums[r] {
			l = m + 1 
			continue
		}

		r = m - 1

		
	}

	return result
}

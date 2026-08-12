func findMin(nums []int) int {
	l, r, result := 0, len(nums)-1, nums[0]

	for l <= r {
		m := (r + l) / 2

		result = min(result, nums[m])

		switch {
			case nums[m] > nums[r]: l = m + 1
			default: r = m - 1
		}
	}

	return result
}

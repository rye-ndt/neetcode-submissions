func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := l + (r - l) / 2

		if nums[m] == target {
			return m
		}

		if nums[m] > target {
			r = m - 1
			continue
		}

		if nums[m] < target {
			l = m + 1
			continue
		}
	}

	return -1
}

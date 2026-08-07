func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := l + (r - l) / 2

		if nums[m] == target {
			return m
		}

		if nums[l] <= nums[m] {
			// check if the target in this 
			if nums[l] <= target && target < nums[m] {
				// select this half 
				r = m - 1
				continue
			}

			// else, select the other half 
			l = m + 1
			continue
		}

		if nums[m] <= nums[r] {
			if nums[m] < target && target <= nums[r] {
				l = m + 1
				continue
			}

			r = m - 1
			continue
		}
	}

	return -1
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := (l + r) / 2

		switch {
			case nums[m] == target: return m
			case nums[l] < nums[m] || nums[r] < nums[m]: 
				if target >= nums[l] && target < nums[m] {
					r = m - 1
				} else {
					l = m + 1
				}
			default:
				if target > nums[m] && target <= nums[r] {
					l = m + 1
				} else {
					r = m - 1
				}
		}
	}

	return -1
}

func maxArea(nums []int) int {	
	result := 0
	l, r := 0, len(nums)-1

	for l < r {
		result = max(result, (r - l) * min(nums[l], nums[r]))

		if min(nums[l], nums[r]) == nums[l] {
			l++
		} else {
			r--
		}
	}

	return result
}

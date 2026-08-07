func maxArea(nums []int) int {	
	result := 0

	l, r := 0, len(nums)-1

	for l < r {
		w := r-l
		h := min(nums[l], nums[r])

		if h * w > result {
			result = h * w
		}

		if h == nums[l] {
			l++
		} else {
			r--
		}
	}

	return result
}

func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1
	min := nums[l]

	for l < r { // no =
		m := l + (r - l) / 2

		if nums[l] < nums[m] { // this half is sorted
			// then min is l 
			// test the other half to see any smaller 
			if min > nums[l] {
				min = nums[l]
			}

			l = m
			continue
		}

		if nums[m] < nums[r] {
			if min > nums[m] {
				min = nums[m]
			}

			r = m
			continue
		}

		// special case: 2 elements left [5, 1]
		if nums[r] < min {
			min = nums[r]
		}

		break
	}

	return min
}

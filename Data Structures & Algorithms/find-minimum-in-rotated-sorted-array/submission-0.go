func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1
	min := nums[0]

	for l < r {
		m := l + (r - l) / 2

		if nums[l] < nums[m] {
			// this part is sorted 
			// the min is l 
			if nums[l] < min {
				min = nums[l]
			}

			// select the other one
			l = m 
			continue
		}

		if nums[m] < nums[r] {
			// this is sorted 
			// the min is m 
			if nums[m] < min {
				min = nums[m]	
			}

			// select this 
			r = m 
			continue
		}

		if nums[r] < min {
			min = nums[r]
		}

		break
	}

	return min
}

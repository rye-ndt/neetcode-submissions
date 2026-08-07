func search(nums []int, target int) int {
	l, r := 0, len(nums) - 1

	for l <= r {
		m := l + (r - l) / 2

		if nums[m] == target {
			return m // edge
		}

		if nums[l] <= nums[m] {
			if nums[l] <= target && target < nums[m] { // notice the = 	
				// in this half, so select it
				r = m-1
				continue
			} 

			// everything else, select the half right 
			l = m+1
			continue
		}

		if nums[m] < nums[r] {
			if nums[m] < target && target <= nums[r] {
				l = m+1
				continue
			}

			r = m-1
			continue
		}
	}

	return -1
}

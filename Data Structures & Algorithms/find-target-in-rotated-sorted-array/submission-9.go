func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := (l + r) / 2

		fmt.Println("l, m, r", l, m, r)

		if nums[m] == target {
			return m
		}

		if l == r && nums[m] != target {
			return -1
		}

		switch {
			case nums[l] < nums[m]: // 3,4,   5   ,6,1,2
				fmt.Println("went to case 0 ")

				if target >= nums[l] && target < nums[m] {
					r = m - 1
				} else {
					l = m + 1
				}
			case nums[l] > nums[m]: // 6,5,  1  ,2,3,4
				fmt.Println("went to case 1")
				if target > nums[m] && target <= nums[r] {
					l = m + 1
				} else {
					r = m - 1
				}
			case nums[r] < nums[m]: // 3,5
				fmt.Println("went to case 2")
				if target >= nums[l] && target < nums[m] {
					r = m - 1
				} else {
					l = m + 1
				}
			case nums[r] > nums[m]:	// 3,5
				fmt.Println("went to case 3")
				if target > nums[m] && target <= nums[r] {
					l = m + 1
				} else {
					r = m - 1
				}
		}
	}

	return -1
}

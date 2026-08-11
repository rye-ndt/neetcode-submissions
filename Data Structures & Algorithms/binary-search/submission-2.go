func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := (r + l) / 2

		fmt.Println("m", m)

		switch {
			case nums[m] > target: r = m - 1
			case nums[m] < target: l = m + 1
			default: return m
		}
	}

	return -1
}

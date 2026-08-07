func twoSum(nums []int, target int) []int {
	l, r := 0, len(nums)-1

	for l < r {
		total := nums[l] + nums[r]

		if total > target {
			r--
		} else if total < target {
			l++
		} else {
			return []int{l+1, r+1}
		}
	}

	return []int{}
}

func twoSum(nums []int, target int) []int {
	l, r := 0, len(nums)-1

	for l < r {
		switch {
			case nums[l] + nums[r] > target: r--
			case nums[l] + nums[r] < target: l++
			default: return []int{l+1, r+1}
		}
	}

	return []int{}
}

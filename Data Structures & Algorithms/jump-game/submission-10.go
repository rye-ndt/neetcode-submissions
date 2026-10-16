func canJump(nums []int) bool {
	reach := 0

	for i, n := range nums {
		if i > reach { return false }
		reach = max(reach, i + n)
	}

	return true
}
func canJump(nums []int) bool {
	reach := 0

	for i := 0; i < len(nums); i++ {
		if i > reach {
			return false
		}

		reach = max(reach, i + nums[i])
	}

	return true
}

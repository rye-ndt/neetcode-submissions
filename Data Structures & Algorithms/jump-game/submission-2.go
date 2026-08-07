func canJump(nums []int) bool {
	reach := 0

	for i := 0; i < len(nums); i++ {
		if i > reach {
			return false // cannot reach here
		}

		reach = max(reach, nums[i] + i)
	}

	return true
}

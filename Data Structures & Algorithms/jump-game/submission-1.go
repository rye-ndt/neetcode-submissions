func canJump(nums []int) bool {
	// how far can i reach so far?
	reach := 0

	for i := 0; i < len(nums); i++ {
		if i > reach {
			return false // we simply cannot get here
		}

		reach = max(reach, i + nums[i])
	}

	return true
}

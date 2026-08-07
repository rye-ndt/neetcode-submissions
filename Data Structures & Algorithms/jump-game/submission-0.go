func canJump(nums []int) bool {
	// how far can i reach so far?
	reach := 0

	for i := 0; i < len(nums); i++ {
		if reach < i {
			return false 
		}
		
		reach = max(reach, i + nums[i])
	}

	return true
}

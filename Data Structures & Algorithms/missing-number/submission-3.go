func missingNumber(nums []int) int {
	xor := len(nums)

	for i, n := range nums {
		xor ^= n ^ i
	}

	return xor
}
func missingNumber(nums []int) int {
	xor := 0

	for i := 0; i <= len(nums); i++ {
		xor ^= i
	}

	for _, n := range nums {
		xor ^= n
	}

	return xor
}
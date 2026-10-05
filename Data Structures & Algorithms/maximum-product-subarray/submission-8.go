func maxProduct(nums []int) int {
	total, biggest, smallest := nums[0], nums[0], nums[0]

	for _, n := range nums[1:] {
		biggest, smallest = max(n, biggest * n, smallest * n), min(n, biggest * n, smallest * n)

		total = max(biggest, total)
	}

	return total
}
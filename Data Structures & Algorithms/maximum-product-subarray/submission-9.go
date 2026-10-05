func maxProduct(nums []int) int {
	total, biggest, smallest := nums[0], nums[0], nums[0]

	for _, n := range nums[1:] {
		nb, nm := biggest * n, smallest * n
		biggest, smallest = max(n, nb, nm), min(n, nb, nm)

		total = max(biggest, total)
	}

	return total
}
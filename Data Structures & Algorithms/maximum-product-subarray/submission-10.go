func maxProduct(nums []int) int {
	total, highest, lowest := nums[0], nums[0], nums[0]

	for _, n := range nums[1:] {
		nb, nm := highest * n, lowest * n
		highest, lowest = max(n, nb, nm), min(n, nb, nm)

		total = max(highest, total)
	}

	return total
}
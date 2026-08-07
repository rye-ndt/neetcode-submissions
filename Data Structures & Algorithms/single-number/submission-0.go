func singleNumber(nums []int) int {
	seen := map[int]int{}

	for _, n := range nums {
		seen[n]++
	}

	for k, v := range seen {
		if v == 1 {
			return k
		}
	}

	return -1
}

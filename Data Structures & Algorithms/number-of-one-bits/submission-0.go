func hammingWeight(n int) int {
	total := 0
	for _, r := range strconv.FormatInt(int64(n), 2) {
		if r == '1' { total++ }
	}
	return total
}

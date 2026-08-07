func countBits(n int) []int {
	result := make([]int, n+1)

	for i := 0; i <= n; i++ {
		s := strconv.FormatInt(int64(i), 2)

		for _, item := range strings.Split(s, "") {
			if item == "1" {
				result[i]++
			}
		}
	}
	
	return result
}

func lengthOfLongestSubstring(s string) int {
	result := 0
	last := map[rune]int{}
	start := 0

	for i, c := range s {
		if v, found := last[c]; found && v >= start {
			start = v+1
		}

		last[c] = i
		result = max(result, i - start + 1)
	}

	return result
}
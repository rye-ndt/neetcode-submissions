func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	result := 1
	seen := map[byte]int{}
	start := 0

	for i := 0; i < len(s); i++ {
		if index, found := seen[s[i]]; found && index >= start {
			start = index+1
		}

		seen[s[i]] = i

		if i - start + 1 > result {
			result = i - start + 1
		}
	}

	return result
}
func lengthOfLongestSubstring(s string) int {
	result := 0

	lastIndex := map[string]int{}

	cur := 0

	for i, c := range strings.Split(s, "") {
		val, found := lastIndex[c] 
		if !found {
			cur++
		} else {
			cur = min(cur+1, i-val)
		}

		result = max(result, cur)
		lastIndex[c] = i
	}

	return result 
}
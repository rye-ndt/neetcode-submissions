func lengthOfLongestSubstring(s string) int {
	result := 0
	pos := map[rune]int{}
	start := 0

	for curIndex, c := range s {
		if i, found := pos[c]; found && i >= start {
			start = i+1
		}

		pos[c] = curIndex
		result = max(result, curIndex-start+1)
	}

	return result
}
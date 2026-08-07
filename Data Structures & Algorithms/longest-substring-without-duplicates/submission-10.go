func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	result := 1
	seen := map[string]int{}
	chars := strings.Split(s, "")
	start := 0

	fmt.Println("chars ", chars)

	for i := 0; i < len(chars); i++ {
		if index, found := seen[chars[i]]; found && index >= start {
			start = index+1
		}

		seen[chars[i]] = i

		if i - start + 1 > result {
			result = i - start + 1
		}
	}

	return result
}
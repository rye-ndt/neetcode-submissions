func wordBreak(s string, wordDict []string) bool {
	seen := make([]bool, len(s) + 1)
	seen[0] = true

	for i := 1; i <= len(s); i++ {
		for _, w := range wordDict {
			if len(w) > i { continue } // as it cannot fit 

			startPos := i - len(w)

			if seen[startPos] && s[startPos:i] == w {
				seen[i] = true
				break
			}
		}
	}

	return seen[len(s)]
}

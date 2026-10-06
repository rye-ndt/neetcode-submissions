func wordBreak(s string, wordDict []string) bool {
	walk := make([]bool, len(s) + 1)
	walk[0] = true

	for i := 1; i <= len(s); i++ {
		for _, w := range wordDict {
			if len(w) > i { continue }

			start := i - len(w)

			if walk[start] && s[start:i] == w { 
				walk[i] = true
				break
			}
		}
	}

	return walk[len(s)]
}

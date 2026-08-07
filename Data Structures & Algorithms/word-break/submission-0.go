func wordBreak(s string, wordDict []string) bool {
    n := len(s)
	note := make([]bool, n+1)
	note[n] = true 

	for i := n-1; i >= 0; i-- {
		for _, w := range wordDict {
			totalLen := i + len(w)

			if totalLen <= n && s[i:totalLen] == w && note[totalLen] {
				note[i] = true
				break
			}
		}
	}

	return note[0]
}

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	note := make([]bool, n+1)
	note[0] = true

	for i := 1; i <= n; i++ {
		for _, w := range wordDict {
			l := len(w)

			inRange := i - l >= 0

			if !inRange {
				continue
			}

			matchWord := s[i-l:i] == w 
			prevSolved := note[i-l] == true

			if matchWord && prevSolved {
				note[i] = true
				break
			}
		}
	}

	return note[n]
}

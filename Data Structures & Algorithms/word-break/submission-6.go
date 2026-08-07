func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	note := make([]bool, n+1)
	note[0] = true

	for i := 1; i <= n; i++ {
		// try every word 
		for _, w := range wordDict {
			l := len(w)

			canFit := i - l >= 0

			if !canFit {
				continue
			}

			isThatWord := s[i - l:i] == w 
			prevDone := note[i - l] == true

			if isThatWord && prevDone {
				note[i] = true
				break // no more search for word
			}
		}
	}

	return note[n]
}

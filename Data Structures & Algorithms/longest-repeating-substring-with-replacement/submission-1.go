func characterReplacement(s string, k int) int {
	charIndexToCount := [26]int{}

	l := 0
	r := 0
	result := 0
	maxFreqChar := 0

	for r = 0; r < len(s); r++ {
		curCharIndex := s[r] - 'A'

		// increase the counter 
		charIndexToCount[curCharIndex]++

		// find the most freq char 
		if charIndexToCount[curCharIndex] > maxFreqChar {
			maxFreqChar = charIndexToCount[curCharIndex]
		}

		width := r-l+1

		// check if the current is surpassing the number of k 
		// this is the counterintuitive part, since no loop is needed 
		if width - maxFreqChar > k {
			// deduct the most left char by 1, NO LOOP 
			leftCharIndex := s[l] - 'A'
			charIndexToCount[leftCharIndex]--

			// move left char 
			l++
		}

		// recalculate the width 
		width = r - l + 1

		if width > result {
			result = width
		}
	}

	return result 
}

func characterReplacement(s string, k int) int {
	charIndexCount := [26]int{} 

	l := 0
	r := 0
	result := 0
	maxFreq := 0

	for r < len(s) {
		curIndex := s[r] - 'A'
		charIndexCount[curIndex]++

		// update the freq
		if charIndexCount[curIndex] > maxFreq {
			maxFreq = charIndexCount[curIndex]
		}

		// check if current windows is not valid anymore 
		// this is the counterintuitive part, since no loop needed 
		width := r-l+1

		if width - maxFreq > k { // which means more than k chars need to be replaced 
			// deduct the l count, and increase l to reduce the window 
			lIndex := s[l] - 'A'
			charIndexCount[lIndex]--
			l++
		}

		newWidth := r-l+1

		if newWidth > result {
			result = newWidth
		}

		r++
	}

	return result
}

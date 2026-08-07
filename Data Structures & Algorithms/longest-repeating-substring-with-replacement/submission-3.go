func characterReplacement(s string, k int) int {
	charIndexCount := map[byte]int{} 

	l := 0
	r := 0
	result := 0
	maxFreqChar := 0

	for r < len(s) {
		curIndex := s[r] - 'A' // to calculate the index of char in alphabet
		charIndexCount[curIndex]++

		if charIndexCount[curIndex] > maxFreqChar {
			maxFreqChar = charIndexCount[curIndex]
		}

		width := r-l+1

		if width - maxFreqChar > k {
			// needs replace 
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

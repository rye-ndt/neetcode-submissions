func characterReplacement(s string, k int) int {
	count := [26]int{}

	l := 0
	maxFreq := 0
	result := 0

	for r := 0; r < len(s); r++ {
		curChar := s[r]
		curCharIndex := curChar - 'A' 

		count[curCharIndex]++ // increase the number of the current 

		if count[curCharIndex] > maxFreq {
			maxFreq = count[curCharIndex]
		}

		width := r-l+1

		if width - maxFreq > k { // which means more than k chars need to be replaced
			leftCharIndex := s[l] - 'A'
			count[leftCharIndex]--
			l++
		}

		width = r-l+1 //recalculate since l is updated

		if width > result {
			result = width
		}
	}

	return result 
}

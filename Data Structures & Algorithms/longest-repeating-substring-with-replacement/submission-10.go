func characterReplacement(s string, k int) int {
	result, maxFreq, l, r := 0, 0, 0, 0
	book := map[byte]int{} 

	for r < len(s) {
		book[s[r]]++
		maxFreq = max(book[s[r]], maxFreq)

		// maintain a valid window
		for r - l + 1 - maxFreq > k {
			book[s[l]]--
			l++
		}

		result = max(result, r - l + 1)

		r++
	}

	return result
}

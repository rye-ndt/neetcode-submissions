func maxChar(book map[byte]int) int {
	curMax := 0

	for _, v := range book {
		curMax = max(curMax, v)
	}

	return curMax
}

func characterReplacement(s string, k int) int {
	result, l, r := 0, 0, 0
	book := map[byte]int{} 

	for r < len(s) {
		book[s[r]]++

		for r - l + 1 - maxChar(book) > k {
			book[s[l]]--

			l++
		}

		result = max(result, r - l + 1)

		r++
	}

	return result
}

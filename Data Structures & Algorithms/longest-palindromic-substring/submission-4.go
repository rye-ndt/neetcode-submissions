func longestPalindrome(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		r, l := i, i

		for r + 1 < len(s) && s[r+1] == s[i] {
			r++ 
		}

		for l > 0 && r < len(s) - 1 && s[l-1] == s[r+1] {
			r++
			l--
		}

		if r - l + 1 > len(result) { result = s[l:r + 1]}
	}

	return result
}
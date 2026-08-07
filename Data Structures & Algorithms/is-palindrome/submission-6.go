func isValid(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsNumber(c)
}

func isPalindrome(s string) bool {
	runes := []rune(s)

	l := 0
	r := len(runes) - 1

	for l < r {
		cl := runes[l]
		cr := runes[r]

		if string(cl) == "" || !isValid(cl) {
			l++
			continue
		}

		if string(cr) == "" || !isValid(cr) {
			r--
			continue
		}

		if strings.ToLower(string(cl)) != strings.ToLower(string(cr)) {
			return false
		}

		l++
		r--
	}

	return true
}

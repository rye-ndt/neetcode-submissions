func valid(v byte) bool {
	return (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9')
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	l, r := 0, len(s)-1

	for l < r {
		if string(s[l]) == "" || !valid(s[l]) {
			l++ 
			continue
		}

		if string(s[r]) == "" || !valid(s[r]) {
			r--
			continue
		}

		if s[l] != s[r] {
			return false
		}
		
		l++
		r--
	}

	return true
}

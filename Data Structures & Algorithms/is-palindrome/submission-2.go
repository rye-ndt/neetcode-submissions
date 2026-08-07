func isAlphaNum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func isPalindrome(s string) bool {
	chars := strings.Split(s, "")
	filtered := []string{}

	for _, char := range chars {
		r := []rune(char)[0]

		if isAlphaNum(r){
			filtered = append(filtered, strings.ToLower(char))
		}
	}

	str := strings.Join(filtered, "")
	start := 0
	end := len(str) - 1 

	for start < end {
		if str[start] != str[end] {
			return false
		}

		start++
		end--
	}

	return true
}

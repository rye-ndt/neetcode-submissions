func partition(s string) [][]string {
	result := [][]string{}

	var backtrack func(path []string, start int)
	backtrack = func(path []string, start int) {
		if start == len(s) {
			result = append(result, append([]string{}, path...))
			return 
		}

		for end := start + 1; end <= len(s); end++ {
			sub := s[start:end]
			if isPalin(sub) {
				backtrack(append(path, sub), end)
			}
		}
	}	

	backtrack([]string{}, 0)
	return result
}

func isPalin(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] { return false }
		l++
		r--
	}
	return true 
}
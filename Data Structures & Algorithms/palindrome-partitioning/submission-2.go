func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}

func clone(a []string) []string{
	cl := make([]string, len(a))
	copy(cl, a)
	return cl
}

func partition2(s string) [][]string {
	result := [][]string{}
	path := []string{}

	var backtrack func(start int) 

	backtrack = func(start int) {
		if start == len(s) {
			result = append(result, clone(path))
			return
		}

		for end := start; end < len(s); end++ {
			choice := s[start:end+1] 

			if isPalindrome(choice) {
				path = append(path, choice)
				backtrack(end+1)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(0)

	return result
}

func partition(s string) [][]string {
	result := [][]string{}
	path := []string{}

	var backtrack func(startIndex int)

	backtrack = func(startIndex int) {
		if startIndex == len(s) {
			result = append(result, clone(path))
			return
		}

		for end := startIndex; end < len(s); end++ {
			choice := s[startIndex:end+1]
			
			if isPalindrome(choice) {
				path = append(path, choice)
				backtrack(end+1)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(0)

	return result
}

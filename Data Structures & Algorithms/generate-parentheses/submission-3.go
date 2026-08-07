func generateParenthesis(n int) []string {
	result := []string{}
	var backtrack func(path string, open, close int)

	backtrack = func(path string, open, close int) {
		if len(path) == 2*n {
			result = append(result, path)
			return
		}
		if open < n {
			backtrack(path+"(", open+1, close)
		}
		if close < open {
			backtrack(path+")", open, close+1)
		}
	}

	backtrack("", 0, 0)
	return result
}
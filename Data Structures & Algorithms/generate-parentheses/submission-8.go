func generateParenthesis(n int) []string {
	start := []string{"(", ")"}

	result := []string{}
	sigs := map[string]bool{}

	var backtrack func(path string)
	backtrack = func(path string) {
		if len(path) == n * 2 { 
			if valid(path) && !sigs[path] { 
				result = append(result, path) 
				sigs[path] = true
			}

			return
		}

		for i := 0; i < len(start); i++ {
			backtrack(path + start[i])
		}
	}	

	backtrack("")

	return result
}

func valid(s string) bool {
	stack := 0

	for _, c := range s {
		switch c {
			case '(': stack++
			case ')': stack--
		}
		
		if stack < 0 { return false }
	}

	return stack == 0
}
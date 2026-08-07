func valid(str string) bool {
	openStack := 0

	for _, s := range str {
		if s == '(' {
			openStack += 1
			continue
		}
		
		// )
		openStack -= 1

		if openStack == -1 {
			return false
		}
	}

	return openStack == 0 
}

func generateParenthesis(n int) []string {
	result := []string{}

	choices := []string{"(", ")"}

	var backtrack func(path string) 

	backtrack = func(path string) {
		if len(path) == n*2 {
			if valid(path){
				result = append(result, path)
			}

			return
		}

		for _, c := range choices {
			path = path + c
			backtrack(path)
			ps := strings.Split(path, "") 
			path = strings.Join(ps[:len(ps)-1], "")
		}
	}

	backtrack("")

	return result
}

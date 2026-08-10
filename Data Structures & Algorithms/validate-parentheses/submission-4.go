func isValid(s string) bool {
	stack := []rune{}

	for _, c := range s {
		if c == '[' || c == '{' || c == '(' {
			stack = append(stack, c)
		} else if len(stack) == 0 {
			return false
		} else {
			last := stack[len(stack)-1]

			case1 := last == '[' && c == ']'
			case2 := last == '{' && c == '}'
			case3 := last == '(' && c == ')'

			if !case1 && !case2 && !case3 {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
 		}
	}

	return len(stack) == 0 
}

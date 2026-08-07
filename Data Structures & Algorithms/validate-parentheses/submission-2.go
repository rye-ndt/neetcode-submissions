func isOpen(c string) bool {
	return c == "(" || c == "{" || c == "["
}

func closeCorrect(c, stackChar string) bool {
	return  (c == ")" && stackChar == "(") || 
			(c == "]" && stackChar == "[") || 
			(c == "}" && stackChar == "{")
}

func isValid(s string) bool {
	stack := []string{}

	cs := strings.Split(s, "")

	for _, c := range cs {
		if isOpen(c) {
			stack = append(stack, c)
			continue
		}

		// this is a close param 
		if len(stack) == 0 || !closeCorrect(c, stack[len(stack) - 1]) {
			return false
		}

		stack = stack[:len(stack) - 1]
	}

	return len(stack) == 0
}

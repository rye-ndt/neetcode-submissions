func evalRPN(tokens []string) int {
	vals := []int{}

	for _, t := range tokens {
		if v, err := strconv.Atoi(t); err == nil {
			vals = append(vals, v)
			continue
		}

		a := vals[len(vals)-2]
		b := vals[len(vals)-1]
		vals = vals[:len(vals)-2]
			
		switch t {
			case "+": vals = append(vals, a + b)
			case "-": vals = append(vals, a - b)
			case "*": vals = append(vals, a * b)
			case "/": vals = append(vals, a / b)
		}
	}

	return vals[0]
}

func evalRPN(tokens []string) int {
	store := []int{}

	for _, t := range tokens {
		if val, err := strconv.Atoi(t); err == nil {
			store = append(store, val)
			continue
		}

		// this is an arithmetric operator 

		if len(store) < 2 {
			continue // cannot calculate anything here
		}

		first := store[len(store) - 2]
		sec := store[len(store) - 1]
		val := 0

		switch t {
			case "+": 
				val = first + sec
			case "-":
				val = first - sec
			case "*":
				val = first * sec
			case "/":
				val = first / sec
		}

		store = store[:len(store) - 2]
		store = append(store, val)
	}

	return store[0]
}

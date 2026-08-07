func pop(a []int) []int {
	return a[:len(a) - 1]
}

func top(a []int) int {
	return a[len(a) - 1]
}

/*
	- maintain a decreasing unresolved list (stack)
	- at each step, loop to see if this current element can solve any item of the stack 
	- pop those elements 
	- append the current to the stack
	- continue 
	remember to store indexes, not values, 
	or else we cannot calculate the result the problem wants
*/

func dailyTemperatures(temps []int) []int {
	pendingIndex := []int{}
	result := make([]int, len(temps)) // init everything to 0

	for i := 0; i < len(temps); i++ {
		cur := temps[i]

		for len(pendingIndex) > 0 && cur > temps[top(pendingIndex)] {
			topIndex := top(pendingIndex)

			resolvedAt := i - topIndex // the distance 
			result[topIndex] = resolvedAt

			// pop 
			pendingIndex = pop(pendingIndex)
		}

		// append the current to the stack 
		pendingIndex = append(pendingIndex, i)
	}

	return result
}

// mental model: dont waste the loop work 
// store the temps as a decreasing stack (must be decreasing to avoid wasting anything)

func top(a []int) int {
	return a[len(a) - 1]
}

func pop (a []int) []int { 
	return a[:len(a) - 1]
}

func dailyTemperatures(temps []int) []int {
	unresolvedIndex := []int{}
	result := make([]int, len(temps))

	for i := 0; i < len(temps); i++ {
		cur := temps[i]

		// this loop enforces the decreasing stack 
		for len(unresolvedIndex) > 0 && cur > temps[top(unresolvedIndex)] {
			j := top(unresolvedIndex)
			result[j] = i - j
			unresolvedIndex = pop(unresolvedIndex)
		}

		unresolvedIndex = append(unresolvedIndex, i)
	}

	return result
}

func dailyTemperatures(temps []int) []int {
	pending := []int{}
	result := make([]int, len(temps))

	for i, t := range temps {
		if len(pending) == 0 {
			pending = append(pending, i)
			continue
		}

		for len(pending) > 0 && t > temps[pending[len(pending)-1]] {
			result[pending[len(pending)-1]] = i - pending[len(pending)-1]
		

			pending = pending[:len(pending)-1]
		}

		pending = append(pending, i)
	}

	return result
}

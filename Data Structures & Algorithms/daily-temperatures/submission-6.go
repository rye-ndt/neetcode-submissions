func dailyTemperatures(temps []int) []int {
	p := []int{} // p = pending
	result := make([]int, len(temps))

	for i, t := range temps {
		for len(p) > 0 && t > temps[p[len(p)-1]] {
			result[p[len(p)-1]] = i  - p[len(p)-1]
			p = p[:len(p)-1]
		}

		p = append(p, i)
	}

	return result
}

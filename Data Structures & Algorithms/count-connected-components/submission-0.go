func countComponents(n int, edges [][]int) int {
    result := 0

	mapper := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]

		mapper[a] = append(mapper[a], b)
		mapper[b] = append(mapper[b], a)
	}

	visited := make([]bool, n)

	var traverse func(cur int) 
	traverse = func(cur int) {
		if visited[cur] {
			return
		}

		visited[cur] = true

		for _, d := range mapper[cur] {
			traverse(d)
		}
	}

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}

		traverse(i)
		result++
	}

	return result
}

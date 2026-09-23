func countComponents(n int, edges [][]int) int {
	graph := map[int][]int{}
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	counter := 0 
	seen := make([]bool, n)
	var dfs func(int) 
	dfs = func(cur int) {
		seen[cur] = true
		options := graph[cur]

		for _, o := range options {
			if !seen[o] {
				dfs(o)
			}
		}
	}

	for num, checked := range seen {
		if checked { continue }
		counter++
		dfs(num)
	}

	return counter
}
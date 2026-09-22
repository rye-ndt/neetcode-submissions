func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 { return false }

	graph := map[int][]int{}
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	visited := map[int]bool{}
	
	var dfs func(last int) 
	dfs = func(last int) {
		visited[last] = true

		for _, item := range graph[last] {
			if !visited[item] { dfs(item) }
		}
	}	

	dfs(0)

	return len(visited) == n
}
func findRedundantConnection(edges [][]int) []int {
	n := 0
	graph := map[int]map[int]bool{}

	for _, e := range edges {
		if _, found := graph[e[0]]; !found { 
			graph[e[0]] = map[int]bool{}
		}

		if _, found := graph[e[1]]; !found { 
			graph[e[1]] = map[int]bool{}
		}

		graph[e[0]][e[1]] = true
		graph[e[1]][e[0]] = true
		n = max(n, max(e[0], e[1]))	
	}

	var dfs func(cur int)

	for i := len(edges)-1; i >= 0; i-- {
		delete(graph[edges[i][0]], edges[i][1])
		delete(graph[edges[i][1]], edges[i][0])
		
		walked := map[int]bool{}

		dfs = func(cur int){
			walked[cur] = true
			for option := range graph[cur] {
				if !walked[option] { dfs(option) }
			}
		}

		dfs(1)

		if len(walked) == n { return edges[i] }

		graph[edges[i][0]][edges[i][1]] = true
		graph[edges[i][1]][edges[i][0]] = true
	}

	return []int{}
}
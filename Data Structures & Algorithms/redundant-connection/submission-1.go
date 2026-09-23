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

	for i := len(edges)-1; i >= 0; i-- {
		pair := edges[i]
		fmt.Println("pair", pair, n)

		// fail fast: if a node only have 1 connection
		if len(graph[pair[0]]) == 1 || len(graph[pair[1]]) == 1 { continue }

		delete(graph[pair[0]], pair[1])
		delete(graph[pair[1]], pair[0])
		
		walked := map[int]bool{}

		var dfs func(cur int)
		dfs = func(cur int){
			walked[cur] = true
			for option := range graph[cur] {
				if !walked[option] { dfs(option) }
			}
		}

		dfs(1)

		fmt.Println("walked", walked)

		if len(walked) == n { return pair }

		graph[pair[0]][pair[1]] = true
		graph[pair[1]][pair[0]] = true
	}

	return []int{}
}
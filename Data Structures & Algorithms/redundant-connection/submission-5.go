func findRedundantConnection(edges [][]int) []int {
	n := 0
	g := map[int]map[int]bool{}

	for _, e := range edges {
		if _, found := g[e[0]]; !found { g[e[0]] = map[int]bool{} }
		if _, found := g[e[1]]; !found {  g[e[1]] = map[int]bool{} }
		g[e[0]][e[1]] = true
		g[e[1]][e[0]] = true
		n = max(n, max(e[0], e[1]))	
	}

	var dfs func(cur int)

	for i := len(edges)-1; i >= 0; i-- {
		delete(g[edges[i][0]], edges[i][1])
		delete(g[edges[i][1]], edges[i][0])
		
		w := map[int]bool{}

		dfs = func(cur int){
			w[cur] = true
			for o := range g[cur] { if !w[o] { dfs(o) }	}
		}

		dfs(1)

		if len(w) == n { return edges[i] }

		g[edges[i][0]][edges[i][1]] = true
		g[edges[i][1]][edges[i][0]] = true
	}

	return []int{}
}
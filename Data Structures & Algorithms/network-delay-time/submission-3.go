func networkDelayTime(times [][]int, n int, k int) int {
    graph := map[int]map[int]int{}
	for _, t := range times {
		if _, found := graph[t[0]]; !found {
			graph[t[0]] = map[int]int{}
		}

		graph[t[0]][t[1]] = t[2]
	}

	dist := map[int]int{}

	var dfs func(int, int)  
	dfs = func(cost, cur int) {
		if d, found := dist[cur]; found && d <= cost { return }
		dist[cur] = cost

		for o, oCost := range graph[cur] {
			dfs(cost+oCost, o)
		}
	}

	dfs(0, k)

	if len(dist) < n { 
		return -1 
	} 

	result := 0
	for _, d := range dist { result = max(result, d) }
	return result
}
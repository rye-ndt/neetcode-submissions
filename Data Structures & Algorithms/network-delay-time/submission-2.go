func networkDelayTime(times [][]int, n int, k int) int {
    graph := map[int]map[int]int{}
	for _, t := range times {
		source := t[0]
		dest := t[1]
		time := t[2]

		if _, found := graph[source]; !found {
			graph[source] = map[int]int{}
		}

		graph[source][dest] = time
	}

	dist := map[int]int{}

	var dfs func(cost, cur int)  
	dfs = func(cost, cur int) {
		if d, found := dist[cur]; found && d <= cost { return }
		dist[cur] = cost

		for option, optionCost := range graph[cur] {
			dfs(cost+optionCost, option)
		}
	}

	dfs(0, k)

	if len(dist) < n { return -1 }

	result := 0
	for _, d := range dist { result = max(result, d) }

	return result
}

// tree with n nodes = acyclic + n-1 edges. this is the rule

func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	mapper := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]

		mapper[a] = append(mapper[a], b)
		mapper[b] = append(mapper[b], a)
	}

	// verify: if it can traverse through all node from node 0
	visited := make([]bool, n)

	var traverse func(cur int) int 
	traverse = func(cur int) int {
		if visited[cur] {
			return 0
		}

		visited[cur] = true

		counter := 1 // current node 

		for _, e := range mapper[cur] {
			counter += traverse(e)
		}

		return counter
	}

	traverseLen := traverse(0)

	return traverseLen == n
}

func findRedundantConnection(edges [][]int) []int {
	n := len(edges) + 1

	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	var find func(x int) int 
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}

		return parent[x]
	}

	for _, e := range edges {
		u, v := find(e[0]), find(e[1])
		if u == v {
			return e
		}

		parent[u] = v
	}

	return nil
}

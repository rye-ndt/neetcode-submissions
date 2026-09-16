type node struct {
	val int 
	dependees []*node
}

func canFinish(n int, pre [][]int) bool {
	nodeMap := map[int]*node{}

	for _, p := range pre {
		if _, found := nodeMap[p[1]]; !found {
			nodeMap[p[1]] = &node{ val: p[1] }
		}

		if _, found := nodeMap[p[0]]; !found {
			nodeMap[p[0]] = &node{ val: p[0] }
		}

		nodeMap[p[1]].dependees = append(nodeMap[p[1]].dependees, nodeMap[p[0]])
	}

	path := map[int]bool{}
	done := map[int]bool{}

	var dfs func(cur *node) bool
	dfs = func(cur *node) bool {
		if done[cur.val] { return true }
		if path[cur.val] { return false }

		path[cur.val] = true 

		for _, child := range cur.dependees {
			if !dfs(child) { return false }
		}

		path[cur.val] = false 
		done[cur.val] = true

		return true
	}	

	for _, node := range nodeMap {
		if !dfs(node) { return false}
	}

	return true
}
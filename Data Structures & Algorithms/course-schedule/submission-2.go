// cycle detection 

func canFinish(numCourses int, prerequisites [][]int) bool {
	mapper := make([][]int, numCourses)

	for _, p := range prerequisites {
		a := p[0]
		b := p[1]

		mapper[a] = append(mapper[a], b)
	}

	freezing := make([]bool, numCourses)
	seen := make([]bool, numCourses)

	var dfs func(cur int) bool 

	dfs = func(cur int) bool {
		if seen[cur] && freezing[cur] {
			return false
		}

		if seen[cur] && !freezing[cur] { 
			return true
		}

		freezing[cur] = true
		seen[cur] = true

		for _, p := range mapper[cur] {
			if !dfs(p) {
				return false
			}
		}

		freezing[cur] = false 

		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}

	return true
}

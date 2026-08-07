// cycle detection 

func canFinish(numCourses int, prerequisites [][]int) bool {
	mapper := make([][]int, numCourses)
	
	// so we know the list of prerequisites of a course 
	for _, p := range prerequisites {
		a := p[0]
		b := p[1]

		mapper[a] = append(mapper[a], b)
	}

	tracker := make([]int, numCourses) // 2 = safe, 1 = ongoing 

	var dfs func(cur int) bool 

	dfs = func(cur int) bool {
		if tracker[cur] == 1 {
			return false // because it is ongoing 
		}

		if tracker[cur] == 2 {
			return true
		}

		tracker[cur] = 1

		for _, p := range mapper[cur] {
			if !dfs(p) {
				return false
			}
		}

		tracker[cur] = 2 // no prereq is blocking it

		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}

	return true
}

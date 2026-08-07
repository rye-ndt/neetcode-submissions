func findOrder(numCourses int, prerequisites [][]int) []int {
	result := []int{}
	prereq := make([][]int, numCourses)

	for _, p := range prerequisites {
		a := p[0]
		b := p[1]

		prereq[a] = append(prereq[a], b)
	}	

	freezing := make([]bool, numCourses)
	safe := make([]bool, numCourses)

	var dfs func(cur int) bool 
	dfs = func(cur int) bool {
		if freezing[cur] {
			return false
		}

		if safe[cur] {
			return true
		}

		freezing[cur] = true

		for _, p := range prereq[cur] {
			if !dfs(p) {
				return false
			}
		}

		freezing[cur] = false 
		safe[cur] = true
		result = append(result, cur)
		
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return []int{}
		}
	}

	return result 
}

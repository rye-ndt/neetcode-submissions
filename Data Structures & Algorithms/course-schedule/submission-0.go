// cycle detection 

func canFinish(numCourses int, prerequisites [][]int) bool {
	preReqList := make([][]int, numCourses)

	for _, p := range prerequisites {
		a := p[0]
		b := p[1]
		preReqList[a] = append(preReqList[a], b)
	}

	state := make([]int, numCourses)

	var dfs func(c int) bool 

	dfs = func(cur int) bool {
		if state[cur] == 1 { // seen it
			return false
		}

		if state[cur] == 2 {
			return true
		}

		state[cur] = 1 

		for _, p := range preReqList[cur] {
			if !dfs(p) {
				return false
			}
		}

		state[cur] = 2 
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}

    return true
}

func findOrder(numCourses int, prerequisites [][]int) []int {
    graph := make([][]int, numCourses)
    indeg := make([]int, numCourses)
    for _, p := range prerequisites {
        a, b := p[0], p[1]
        graph[b] = append(graph[b], a)
        indeg[a]++
    }

    queue := []int{}
    for i := 0; i < numCourses; i++ {
        if indeg[i] == 0 {
            queue = append(queue, i)
        }
    }

    order := []int{}
    for len(queue) > 0 {
        c := queue[0]
        queue = queue[1:]
        order = append(order, c)
        for _, nxt := range graph[c] {
            indeg[nxt]--
            if indeg[nxt] == 0 {
                queue = append(queue, nxt)
            }
        }
    }

    if len(order) == numCourses {
        return order
    }
    return []int{}
}
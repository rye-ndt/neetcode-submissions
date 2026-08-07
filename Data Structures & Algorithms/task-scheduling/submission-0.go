type HItem struct {
	Char string
	Count int
}

type H []HItem

func (h H) Len() int {
	return len(h)
}

func (h H) Less(i, j int) bool {
	return h[i].Count > h[j].Count
} 

func (h H) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *H) Push(x any) {
	*h = append(*h, x.(HItem))
}

func (h *H) Pop() any {
	old := *h

	if len(old) == 0 {
		return nil
	}

	last := old[len(old) - 1]
	*h = old[:len(old) - 1]

	return last
}

func leastInterval(tasks []byte, n int) int {
	h := &H{}
	heap.Init(h)

	seen := map[string]int{}

	for _, t := range tasks {
		seen[string(t)]++
	}

	for k, v := range seen {
		heap.Push(h, HItem{
			Char: k,
			Count: v,
		})
	}

	fmt.Println("heap: ", h)

	cycles := 0

	for h.Len() > 0 {
		newItems := []HItem{}

		for i := 0; i <= n; i++ {
			if h.Len() > 0 {
				cur := heap.Pop(h).(HItem)

				if cur.Count > 1 {
					newItems = append(newItems, HItem{
						Char: cur.Char,
						Count: cur.Count-1,
					})
				}

				cycles++
			} else if len(newItems) > 0 {
				cycles++
			}

			if h.Len() == 0 && len(newItems) == 0 {
				break
			}
		}

		for _, t := range newItems {
			heap.Push(h, t)
		}
	}

	return cycles
}

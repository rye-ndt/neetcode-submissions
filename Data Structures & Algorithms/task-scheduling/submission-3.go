type HItem struct {
	v string 
	counter int
	cooldown int
}

type H []*HItem

func (h H) Len() int { return len(h)}
func (h H) Less(i, j int) bool { 
	iReady, jReady := h[i].cooldown <= 0, h[j].cooldown <= 0
	if iReady != jReady {
		return iReady
	}

	return h[i].counter > h[j].counter
}

func (h H) Swap(i, j int) {h[i], h[j] = h[j], h[i] }
func (h *H) Push(v any) { *h = append(*h, v.(*HItem)) }
func (h *H) Pop() any {
	clone := *h 
	last := clone[len(clone)-1]
	*h = clone[:len(clone)-1]
	return last
} 

func leastInterval(tasks []byte, n int) int {
	store := map[string]int{}
	for _, t := range tasks {
		store[string(t)]++
	}

	h := &H{}
	for v, counter := range store {
		heap.Push(h, &HItem{
			v: v,
			counter: counter,
			cooldown: 0,
		})
	}

	stack := []string{}

	for h.Len() > 0 {
		item := heap.Pop(h).(*HItem)

		if item.cooldown <= 0 {
			stack = append(stack, item.v)
			item.counter--
			item.cooldown = n + 1
		} else {
			stack = append(stack, "nil")
		}

		for _, hi := range *h {
			hi.cooldown--
		}
		
		item.cooldown--
		heap.Init(h)

		if item.counter > 0 {
			heap.Push(h, item)
		}
	}

	return len(stack)
}
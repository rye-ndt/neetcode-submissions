type HItem struct {
	v string 
	counter int
	availableAt int
}

type H []*HItem

var time = 0

func (h *HItem) ready() bool {
	return h.availableAt <= time 
}

func (h H) Len() int { return len(h)}

func (h H) Less(i, j int) bool { 
	if h[i].ready() != h[j].ready() {
		return h[i].ready()
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
	store := map[byte]int{}
	for _, t := range tasks {
		store[t]++
	}

	h := &H{} // ready items only — Less is just counter descending
	for v, c := range store {
		heap.Push(h, &HItem{v: string(v), counter: c})
	}

	type cooling struct {
		item        *HItem
		availableAt int
	}
	queue := []cooling{} // FIFO, naturally sorted by availableAt

	t := 0
	for h.Len() > 0 || len(queue) > 0 {
		if len(queue) > 0 && queue[0].availableAt <= t {
			heap.Push(h, queue[0].item)
			queue = queue[1:]
		}

		if h.Len() > 0 {
			item := heap.Pop(h).(*HItem)
			item.counter--
			if item.counter > 0 {
				queue = append(queue, cooling{item, t + n + 1})
			}
		}
		t++
	}
	return t
}
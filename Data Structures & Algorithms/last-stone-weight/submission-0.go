
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h 
	last := old[old.Len() - 1]

	*h = old[:old.Len() - 1]

	return last
}

func lastStoneWeight(stones []int) int {
	h := &IntHeap{}

	heap.Init(h)

	for _, n := range stones {
		heap.Push(h, n)
	}

	for h.Len() > 0 {
		top1 := heap.Pop(h).(int)

		if h.Len() == 0 {
			return top1
		}

		top2 := heap.Pop(h).(int)

		heap.Push(h, abs(top1 - top2))
	}

	return 0
}

func abs(x int) int {
	if x > 0 {
		return x
	}

	return -x
}
type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(v any) { *h = append(*h, v.(int)) }
func (h *IntHeap) Pop() any {
	clone := *h 
	last := clone[len(clone)-1]
	*h = clone[:len(clone)-1]

	return last
}

func lastStoneWeight(stones []int) int {
	h := &IntHeap{}

	for _, s := range stones {
		heap.Push(h, s)
	}

	for h.Len() > 1 {
		a, b := heap.Pop(h).(int), heap.Pop(h).(int)
		heap.Push(h, max(a - b, b - a))
	}

	return h.Pop().(int)
}
type H []int

func (h H) Len() int {
	return len(h)
}

func (h H) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h H) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *H) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *H) Pop() any {
	old := *h

	last := old[old.Len() - 1]
	*h = old[:old.Len() - 1]

	return last
}

func findKthLargest(nums []int, k int) int {
	h := &H{}
	heap.Init(h)

	for _, n := range nums {
		heap.Push(h, n)

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return heap.Pop(h).(int)
}

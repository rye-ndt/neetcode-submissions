type IntHeap []int
func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(v any) {*h = append(*h, v.(int))}
func (h *IntHeap) Pop() any {
	clone := *h
	last := clone[len(clone)-1]
	*h = clone[:len(clone)-1]
	return last
}

func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}

	for _, n := range nums {
		heap.Push(h, n)
	}

	result := 0

	for i := 0; i < k; i++ {
		result = heap.Pop(h).(int)
	}

	return result 
}

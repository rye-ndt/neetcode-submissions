type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(ele any) { *h = append(*h, ele.(int)) }
func (h *IntHeap) Pop() any {
	clone := *h 
	last := clone[len(clone)-1]
	*h = clone[:len(clone)-1]
	return last
}

type KthLargest struct {
	k int
	h *IntHeap
}

func Constructor(k int, nums []int) KthLargest {
	h := &IntHeap{}

	for _, n := range nums {
		heap.Push(h, n)
	}

	return KthLargest{
		k: k,
		h: h,
	}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.h, val)

	result := 0
	popped := []int{}

	for i := 1; i <= this.k; i++ {
		result = heap.Pop(this.h).(int)
		popped = append(popped, result)
	}

	for _, p := range popped {
		heap.Push(this.h, p)
	}

	return result
}

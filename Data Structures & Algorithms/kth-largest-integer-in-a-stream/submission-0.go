// what is a heap?
// - a tree structure
// - the root is the smallest (or biggest) element 
// - the rest are unsorted. the heap only ensure the top is either smallest or biggest
// - handle the stream / continuous problems that sorting is expensive

// mental model:
// - min or max heap?
// - what do i push?
// - when do i trim?

// in golang, you must define a boilerplate

type IntHeap []int 

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j] // min heap
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h

	lastEle := old[len(old)-1]
	*h = old[:len(old)-1]

	return lastEle
}

type KthLargest struct {
    k int
	h *IntHeap
}


func Constructor(k int, nums []int) KthLargest {
    h := &IntHeap{}
	heap.Init(h)

	kl := KthLargest{
		k: k,
		h: h,
	}

	for _, n := range nums {
		kl.Add(n)
	}

	return kl
}


func (this *KthLargest) Add(val int) int {
	heap.Push(this.h, val)

	if this.h.Len() > this.k {
		heap.Pop(this.h)
	}

	return (*this.h)[0]
}

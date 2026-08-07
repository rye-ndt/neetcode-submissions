// the best solution is bucket sort, but i dont understand it 
// so use simple algorithm here

type Pair struct {
	Num int
	Count int
}

type Item struct {
	Val int
	Count int
}

type H []*Item

func (h H) Len() int {
	return len(h)
}

func (h H) Less(i, j int) bool {
	return h[i].Count < h[j].Count
} 

func (h H) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *H) Push(x any) {
	*h = append(*h, x.(*Item))
}

func (h *H) Pop() any {
	old := *h
	last := old[len(old) - 1]
	*h = old[:len(old) - 1]

	return last
}

func topKFrequent(nums []int, k int) []int {
	h := &H{}
	heap.Init(h)

	store := map[int]int{}

	for _, n := range nums {
		store[n]++
	}

	for num, count := range store {
		heap.Push(h, &Item{Val: num, Count: count})
		
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := []int{}

	for h.Len() > 0 {
		result = append(result, heap.Pop(h).(*Item).Val)
	}

	return result
}

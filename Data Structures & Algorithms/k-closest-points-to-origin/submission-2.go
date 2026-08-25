type IntHeap []float64 
func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(v any) { *h = append(*h, v.(float64)) }
func (h *IntHeap) Pop() any {
	clone := *h
	last := clone[len(clone)-1]
	*h = clone[:len(clone)-1]
	return last
}

func kClosest(points [][]int, k int) [][]int {
	store := map[float64][][]int{}
	h := &IntHeap{}

	for _, p := range points {
		distance := math.Sqrt(float64(p[0]*p[0] + p[1]*p[1]))

		if _, found := store[distance]; !found {
			store[distance] = [][]int{}
		}
		store[distance] = append(store[distance], p)

		heap.Push(h, distance)
	}

	result := [][]int{}

	for i := 0; i < k; i++ {
		distance := heap.Pop(h).(float64)
		result = append(result, store[distance]...)
		delete(store, distance)
	}

	return result
}

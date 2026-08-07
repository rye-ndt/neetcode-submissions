type Point struct {
	X int
	Y int
	Distance float64
}

type PHeap []*Point

func (h PHeap) Len() int {
	return len(h)
}

func (h PHeap) Less(i, j int) bool {
	return h[i].Distance > h[j].Distance
}

func (h PHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PHeap) Push(x any) {
	*h = append(*h, x.(*Point))
}

func (h *PHeap) Pop() any {
	old := *h
	last := old[old.Len() - 1]
	*h = old[:old.Len() - 1]

	return last
}

func kClosest(points [][]int, k int) [][]int {
	h := &PHeap{}
	heap.Init(h)

	for _, p := range points {
		heap.Push(h, &Point{
			X: p[0],
			Y: p[1],
			Distance: calcDistance(p[0], p[1]),
		})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := [][]int{}

	for _, item := range *h {
		result = append(result, []int{item.X, item.Y})
	}

	return result
}

func calcDistance(x, y int) float64 {
	return math.Hypot(float64(x), float64(y))
}

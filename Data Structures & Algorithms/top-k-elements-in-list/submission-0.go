// the best solution is bucket sort, but i dont understand it 
// so use simple algorithm here

type Pair struct {
	Num int
	Count int
}

func topKFrequent(nums []int, k int) []int {
	freq := map[int]int{}

	for _, num := range nums {
		freq[num]++
	}

	pairs := []Pair{}

	for num, count := range freq {
		pairs = append(pairs, Pair{
			Num: num,
			Count: count,
		})
	}

	// then sort the pair with custom algorithm
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Count > pairs[j].Count
	})

	result := []int{}

	for i := 0; i < k; i++ {
		result = append(result, pairs[i].Num)
	}

	return result
}

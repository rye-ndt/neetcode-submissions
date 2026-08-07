func longestConsecutive(nums []int) int {
	store := map[int]bool{}

	for _, num := range nums {
		store[num] = true
	}

	longest := 0

	for num := range store {
		if store[num-1] {
			continue // ignore the numbers that are not the start off a sequence
		}

		streak := 1
		cur := num

		for store[cur+1] {
			streak++
			cur++
		}

		if streak > longest {
			longest = streak
		}
	}

	return longest
}

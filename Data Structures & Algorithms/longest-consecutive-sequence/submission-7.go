func longestConsecutive(nums []int) int {
	store := map[int]bool{}

	for _, num := range nums {
		store[num] = true
	}

	longest := 0

	for num := range store {
		// only check for elements that are start of a sequence 
		if store[num-1] {
			continue // because previous element exists 
		}

		// check the sequence from now on
		streak := 0
		curNum := num

		for store[curNum] {
			streak++
			curNum++ // loop to the next element
		}

		if streak > longest {
			longest = streak
		}
	}

	return longest
}

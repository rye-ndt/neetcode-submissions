func longestConsecutive(nums []int) int {
	sort.Ints(nums)

	filtered := []int{}

	for _, n := range nums {
		if len(filtered) == 0 || n != filtered[len(filtered)-1] {
			filtered = append(filtered, n)
		}
	}

	longest := 0

	for i := 0; i < len(filtered); i++ {
		curCount := 1

		cur := filtered[i]

		j := i+1

		for j < len(filtered) {
			if filtered[j] != cur+1 {
				break
			} 
			
			curCount++
			cur = filtered[j]
			j++
		}

		if curCount > longest {
			longest = curCount
		}

		i = j-1
	} 

	return longest
}

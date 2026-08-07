func canPartition(nums []int) bool {
	total := 0

	for _, n := range nums {
		total += n
	}

	if total % 2 != 0 {
		return false 
	}

	target := total / 2

	note := make([]bool, target+1)
	note[0] = true

	for _, n := range nums {
		for i := target; i >= n; i-- {
			if note[i-n] {
				note[i] = true
			}
		}
	}

	return note[target]
}

// frame the problem like this:
// at each index, look at the number, jump to that number
// the number to jump to is always deterministic, 
// so fast and slow will eventually meet again 

// i dont understand this problem

func findDuplicate(nums []int) int {
	s := nums[0]
	f := nums[0]

	// find the starting point
	for {
		s = nums[s]
		f = nums[nums[f]]

		if s == f {
			break
		}
	}

	// reset and start with the same speed
	s = nums[0]

	for s != f {
		s = nums[s]
		f = nums[f]
	}

	return f
}

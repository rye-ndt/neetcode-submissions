func twoSum(numbers []int, target int) []int {
	s := 0
	e := len(numbers) - 1

	for s < e {
		cur := numbers[s] + numbers[e]

		if cur > target {
			e--
			continue
		}

		if cur < target {
			s++
			continue
		}

		return []int{1+s, 1+e}
	}

	return []int{}
}

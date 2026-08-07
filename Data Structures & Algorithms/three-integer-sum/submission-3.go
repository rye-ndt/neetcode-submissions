func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	dedup := map[string]bool{}
	result := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		cur := nums[i]

		l := i+1
		r := len(nums) - 1

		for l < r {
			nl := nums[l]
			nr := nums[r]

			total := cur + nr + nl 

			if total > 0 {
				r--
				continue
			}

			if total < 0 {
				l++
				continue
			}

			sig := strconv.Itoa(cur) + strconv.Itoa(nums[l]) + strconv.Itoa(nums[r])

			if !dedup[sig] {	
				result = append(result, []int{cur, nums[l], nums[r]})
				dedup[sig] = true
			}

			l++
			r--
		}
	}

	return result
}

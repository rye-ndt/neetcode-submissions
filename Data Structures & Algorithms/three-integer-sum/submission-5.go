func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	fmt.Println("nums", nums)

	seen := map[[3]int]bool{}

	for i := 0; i < len(nums); i++ {
		cur := nums[i]

		l, r := i+1, len(nums)-1

		for l < r {
			total := cur + nums[l] + nums[r]

			if total > 0 {
				r--
			} else if total < 0 {
				l++
			} else {
				seen[[3]int{cur, nums[l], nums[r]}] = true
				l++
				r--
			}
		}
	}

	result := [][]int{}

	for k, _ := range seen {
		result = append(result, []int{k[0], k[1], k[2]})
	}

	return result
}

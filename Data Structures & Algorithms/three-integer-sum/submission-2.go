func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	// [-3,-1,0,1,2]
	// fix one number position, then find the other 2 

	result := [][]int{}
	dedupSig := map[[3]int]bool{}

	for i := 0; i < len(nums) - 2; i++ {
		mid := nums[i]

		l := i+1
		r := len(nums) - 1

		for l < r {
			sum := nums[l] + nums[r] + mid

			if sum == 0 {
				sig := [3]int{mid, nums[l], nums[r]}

				if _, found := dedupSig[sig]; !found {
					result = append(result, []int{mid, nums[l], nums[r]})
					dedupSig[sig] = true
				} 

				l++
				r--
				continue
			}

			if sum > 0 {
				r--
				continue
			}

			if sum < 0 {
				l++
				continue
			}
		}
	}

	return result
}

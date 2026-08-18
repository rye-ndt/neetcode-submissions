func findDuplicate(nums []int) int {
	low := 1
	high := len(nums)-1

	for low <= high {
		mid := (low + high) / 2

		fmt.Println("lte", mid)

		counter := 0

		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid { counter++ }
		}

		fmt.Println("total", counter)

		if counter > mid {
			fmt.Println("duplication less than mid")
			high = mid-1
		} else {
			fmt.Println("dup bigger than mid")
			low = mid+1
		}
	}

	return low
}

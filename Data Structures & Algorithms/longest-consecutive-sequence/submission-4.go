func longestConsecutive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	sort.Ints(nums) // 0,1,3,4,5,6...

	fmt.Println("nums: ", nums)

	streak := 1
	curStreak := 1

	for i := 1; i < len(nums); i++ {
		if nums[i-1]+1 == nums[i] {
			fmt.Println("at index: ", i)
			curStreak++
		} else if nums[i] == nums[i-1] {
			continue
		} else {
			fmt.Println("reset at index: ", i)
			curStreak = 1
		}

		fmt.Println("cur - streak: ", curStreak, streak)

		if curStreak > streak {
			streak = curStreak
		}
	}

	return streak
}

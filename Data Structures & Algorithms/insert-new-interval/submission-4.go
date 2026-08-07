func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	newStart := newInterval[0]
	newEnd := newInterval[1]

	appended := false

    for i := 0; i < len(intervals); i++ {
		curStart := intervals[i][0]
		curEnd := intervals[i][1]

		if curEnd < newStart {
			continue
		}

		if curStart > newEnd {
			intervals = append(intervals, nil)    
			copy(intervals[i+1:], intervals[i:])
			intervals[i] = newInterval
			appended = true
			break
		} 

		intervals[i] = []int{
			min(curStart, newStart), 
			max(curEnd, newEnd),
		}

		appended = true
	}

	if !appended {
		intervals = append(intervals, newInterval)
	}

	result := [][]int{}

	i := 0

	for i < len(intervals) {
		j := i+1

		curStart := intervals[i][0]
		curEnd := intervals[i][1]

		item := []int{curStart, curEnd}

		fmt.Println("cur item: ", item)

		// [1, 5] [2, 6] [3, 7]
		// [1, 6] [3, 7]
		// [1, 7]

		for j < len(intervals) {
			nextStart := intervals[j][0]
			nextEnd := intervals[j][1]

			if nextStart > curEnd {
				j--
				break
			}

			if curEnd > nextStart {
				item = []int{
					min(item[0], nextStart),
					max(item[1], nextEnd),
				}

				fmt.Println("at j ", j, " item became ", item)
			}

			j++
		}

		result = append(result, item)
		i = j+1
	}

	return result
}

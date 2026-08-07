func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

    for i := 0; i < len(intervals); i++ {
		if intervals[i][1] < newInterval[0] {
			if i == len(intervals) - 1 {
				intervals = append(intervals, newInterval)
			}

			continue
		}

		if intervals[i][0] > newInterval[1] {
			intervals = append(intervals, nil)    
			copy(intervals[i+1:], intervals[i:])
			intervals[i] = newInterval
			break
		} 

		intervals[i] = []int{
			min(intervals[i][0], newInterval[0]), 
			max(intervals[i][1], newInterval[1]),
		}
	}

	result := [][]int{}

	i := 0

	for i < len(intervals) {
		item := []int{intervals[i][0], intervals[i][1]}

		j := i+1

		for j = j; j < len(intervals); j++ {
			if intervals[j][0] > intervals[i][1] {
				j--
				break
			}

			if intervals[i][1] > intervals[j][0] {
				item = []int{
					min(item[0], intervals[j][0]),
					max(item[1], intervals[j][1]),
				}
			}
		}

		result = append(result, item)
		i = j+1
	}

	return result
}

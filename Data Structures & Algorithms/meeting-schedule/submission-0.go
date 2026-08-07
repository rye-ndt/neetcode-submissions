/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	fmt.Println("start with: ", intervals)

	for i := 0; i < len(intervals)-1; i++ {
		j := i+1

		if intervals[j].start < intervals[i].end {
			return false
		}
	}

	return true
}

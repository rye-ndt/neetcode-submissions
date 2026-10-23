/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
	if len(intervals) <= 1 { return true }
	
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	end := intervals[0].end

	for i := 1; i < len(intervals); i++ {
		if intervals[i].start < end { return false }
		end = max(end, intervals[i].end)
	}

	return true
}
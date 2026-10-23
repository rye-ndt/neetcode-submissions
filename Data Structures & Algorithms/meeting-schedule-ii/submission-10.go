/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func minMeetingRooms(nums []Interval) int {
	starts, ends := make([]int, len(nums)), make([]int, len(nums))

	for i := range nums {
		starts[i], ends[i] = nums[i].start, nums[i].end
	}

	sort.Ints(starts)
	sort.Ints(ends)

	room, j := 0, 0 

	for i := 0; i < len(nums); i++ {
		if starts[i] < ends[j] {
			room++
		} else {
			j++
		}
	}

	return room
}
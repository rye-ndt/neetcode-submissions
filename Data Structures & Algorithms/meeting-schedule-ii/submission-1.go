/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func overlap(a, b Interval) bool {
	if a.start < b.start {
		return b.start < a.end 
	}

	return a.start < b.end
}

func head(i []Interval) Interval {
	if len(i) == 0 {
		return Interval{
			start: -1,
			end: -1,
		}
	}

	return i[len(i)-1]
}

func minMeetingRooms(nums []Interval) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i].start < nums[j].start
	})

	rooms := make([][]Interval, len(nums))
	counter := 0

	for _, cur := range nums {
		for i, r := range rooms {
			newRoom := len(r) == 0

			if newRoom || !overlap(cur, r[len(r)-1]) {
				rooms[i] = append(rooms[i], cur)

				if newRoom {
					counter++
				}

				break
			}
		}
	}
	return counter
}

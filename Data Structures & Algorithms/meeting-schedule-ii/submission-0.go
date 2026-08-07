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

	fmt.Println("start with: ", nums)

	rooms := make([][]Interval, len(nums))

	for i := 0; i < len(nums); i++ {
		cur := nums[i]

		for j := 0; j < len(rooms); j++ {
			roomHead := head(rooms[j])

			if !overlap(cur, roomHead) {
				rooms[j] = append(rooms[j], cur)
				break
			}
		}
	}

	fmt.Println("rooms: ", rooms)

	counter := 0

	for _, r := range rooms {
		if len(r) > 0 {
			counter++
		}
	}

	return counter
}

type car struct {
	pos int
	sp int
}

func carFleet(target int, position []int, speed []int) int {
	result := 0
	cars := []car{}

	for i, p := range position {
		cars = append(cars, car{
			pos: p,
			sp: speed[i],
		})
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	fleetTime := 0.0

	for _, c := range cars {
		time := float64(target - c.pos) / float64(c.sp)

		if time > fleetTime {
			result++
			fleetTime = time
		}
	}

	return result
}

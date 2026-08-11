type car struct {
	pos int
	sp int
}

func carFleet(target int, position []int, speed []int) int {
	result := 0
	cars := []*car{}

	for i, p := range position {
		cars = append(cars, &car{
			pos: p,
			sp: speed[i],
		})
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	for i := 0; i < len(cars); i++ {
		fleetLead := cars[i]

		if fleetLead == nil { continue }

		result++

		leadTime := float64(target - fleetLead.pos) / float64(fleetLead.sp)

		for j := i+1; j < len(cars); j++ {
			cur := cars[j]

			if cur == nil { continue }

			time := float64(target - cur.pos) / float64(cur.sp)

			if time <= leadTime {
				cars[j] = nil
			}
		}
	}

	return result
}

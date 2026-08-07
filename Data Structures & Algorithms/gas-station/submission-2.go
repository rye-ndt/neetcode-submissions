// if total gas > total cost -> solution exists 
// the start point is the point where the tank hits its lowest

func canCompleteCircuit(gas []int, cost []int) int {
	tank := 0
	minTank := 0
	start := 0

	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i] // net gain 

		if tank < minTank {
			minTank = tank 
			start = i+1
		}
	}

	if tank < 0 {
		return -1
	}

	return start % len(gas) // so we have a circular 
}

// if total gas > total cost -> solution exists 
// the start point is the point where the tank hits its lowest

func canCompleteCircuit(gas []int, cost []int) int {
	tank, minTank := 0, 0
	result := 0

	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i]

		if minTank > tank {
			minTank = tank
			result = i+1
		} 
	}

	if tank < 0 {
		return -1
	}

	return result % len(gas)
}

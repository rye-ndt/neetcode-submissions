func canCompleteCircuit(gas []int, cost []int) int {
	tank, minTank := 0, 0
	result := 0

	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i]

		if minTank > tank {
			minTank = tank
			result = i+1 
			// why: the start point is where the tank hits lowest
		} 
	}

	// start at any position; we should be able to move
	// entirely 
	if tank < 0 {
		return -1
	}

	return result % len(gas) // ensure the circular 
}

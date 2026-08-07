func canCompleteCircuit(gas []int, cost []int) int {
	tank, minTank := 0, 0
	result := 0

	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i] // net gain 

		if minTank > tank {
			minTank = tank
			result = i+1 
			// why: the start point is where the tank hits lowest
		} 
	}

	// net gain must be positive, else cannot move
	if tank < 0 {
		return -1
	}

	return result % len(gas) // ensure the circular 
}

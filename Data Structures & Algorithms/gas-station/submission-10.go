func canCompleteCircuit(gas []int, cost []int) int {
	minTank, tank := 0, 0
	result := 0

	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i] // net

		if tank < minTank {
			minTank = tank
			result = i+1
		}
	}

	if tank < 0 {
		return -1 // no way to get there 
	}

	return result
}

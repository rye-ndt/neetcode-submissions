// if total gas > total cost -> solution exists 
// the start point is the point where the tank hits its lowest

func canCompleteCircuit(gas []int, cost []int) int {
	total := 0
	run := 0
	minRun := 0
	startAfter := 0

	for i := 0; i < len(gas); i++ {
		run += gas[i] - cost[i]
		total = run 

		if run < minRun {
			minRun = run 
			startAfter = i+1
		}
	}

	if total < 0 {
		return -1
	}

	return startAfter % len(gas)
}

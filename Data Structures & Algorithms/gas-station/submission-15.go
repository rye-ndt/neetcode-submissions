func canCompleteCircuit(gas []int, cost []int) int {
	balance, anchor := 0, 0

	for i := 0; i < len(gas); i++ {
		balance += gas[i] - cost[i]
	}

	if balance < 0 { return -1 }
	balance = 0

	for i := 0; i < len(gas); i++ {
		if i == anchor-1 || (anchor == 0 && i == len(gas)-1) { 
			break 
		}
		
		balance += gas[i] - cost[i]

		if balance < 0 {
			balance = 0 
			anchor = i+1
		}

		if i == len(gas)-1 { i = -1 }
	}

	return anchor
}
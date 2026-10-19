func canCompleteCircuit(gas []int, cost []int) int {
	if sum(gas) < sum(cost) { return -1 }

	balance, result := 0, 0

	for i := 0; i < len(gas); i++ {
		if i == result-1 || (result == 0 && i == len(gas)-1) { 
			break 
		}
		
		balance += gas[i] - cost[i]

		if balance < 0 {
			balance = 0 
			result = i+1
		}

		if i == len(gas)-1 { i = -1 }
	}

	return result
}

func sum(nums []int) int {
	result := 0
	for _, n := range nums { result += n }
	return result
}
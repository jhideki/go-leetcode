package main

func canCompleteCircuit(gas []int, cost []int) int {
	maxIndex := 0
	amount := 0
	diff := 0
	total := 0
	for i := 0; i < len(gas); i++ {
		diff = gas[i] - cost[i]
		total += gas[i] - cost[i]
		amount += diff
		if total < 0 {
			total = 0
			maxIndex = i + 1
		}

	}

	if amount < 0 {
		return -1
	} else {
		return maxIndex
	}
}

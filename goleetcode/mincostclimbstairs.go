package main

func minCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)
	for i := len(cost) - 3; i >= 0; i-- {
		cost[i] = min(cost[i+1], cost[i+2])

	}

	return min(cost[0], cost[1])
}

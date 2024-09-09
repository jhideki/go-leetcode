package main

func findTragetSumWays(nums []int, target int) int {
	type Tuple struct {
		i   int
		val int
	}

	var m = make(map[Tuple]int)
	var rec func(int, int) int

	rec = func(i int, curAmount int) int {
		if i == len(nums) {
			if curAmount == target {
				return 1
			} else {
				return 0
			}
		}

		if val, exists := m[Tuple{i, curAmount}]; exists {
			return val
		}

		m[Tuple{i, curAmount}] = rec(i+1, curAmount+nums[i]) + rec(i+1, curAmount-nums[i])
		return m[Tuple{i, curAmount}]
	}
	return rec(0, 0)
}

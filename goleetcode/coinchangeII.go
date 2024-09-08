package main

func change(amount int, coins []int) int {
	type Tuple struct {
		i   int
		val int
	}

	var m = make(map[Tuple]int)
	var rec func(int, int) int

	rec = func(i int, curAmount int) int {
		if curAmount == amount {
			return 1
		}

		if curAmount > amount || i == len(coins) {
			return 0
		}

		if val, exists := m[Tuple{i, curAmount}]; exists {
			return val
		}

		m[Tuple{i, curAmount}] = rec(i, curAmount+coins[i]) + rec(i+1, amount)
		return m[Tuple{i, curAmount}]
	}
	return rec(0, 0)
}

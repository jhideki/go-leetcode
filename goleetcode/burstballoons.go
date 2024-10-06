package main

func maxCoins(nums []int) int {
	nums = append(nums, 1)
	nums = append([]int{1}, nums...)
	type Tuple struct {
		l int
		r int
	}

	var m = make(map[Tuple]int)
	var rec func(int, int) int

	rec = func(l int, r int) int {
		if l > r {
			return 0
		}

		if val, exists := m[Tuple{l, r}]; exists {
			return val
		}

		m[Tuple{l, r}] = 0
		for i := l; i < r+1; i++ {
			coins := nums[l-1] * nums[i] * nums[r+1]
			coins += rec(l, i-1) + rec(i+1, r)
			m[Tuple{l, r}] = max(m[Tuple{l, r}], coins)
		}
		return m[Tuple{l, r}]
	}

	return rec(1, len(nums)-2)
}

package main

func numDistinct(s string, t string) int {
	type Tuple struct {
		i int
		j int
	}

	var m = make(map[Tuple]int)
	var rec func(int, int) int

	rec = func(i int, j int) int {
		if j == len(t) {
			return 1
		}
		if i == len(s) {
			return 1
		}

		if val, exists := m[Tuple{i, j}]; exists {
			return val
		}

		if s[i] == t[j] {
			m[Tuple{i, j}] = rec(i+1, j+1) + rec(i+1, j)
		} else {
			m[Tuple{i, j}] = rec(i+1, j)
		}
		return m[Tuple{i, j}]
	}
	return rec(0, 0)
}

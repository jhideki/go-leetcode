package main

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)

	type Tuple struct {
		l int
		r int
	}

	var mem = make(map[Tuple]bool)

	var rec func(i, j int) bool
	rec = func(i, j int) bool {
		if j == n {
			return i == m
		}
		if val, exists := mem[Tuple{i, j}]; exists {
			return val
		}
		match := i < m && (s[i] == p[j] || p[j] == '.')
		if j+1 < n && p[j+1] == '*' {
			mem[Tuple{i, j}] = rec(i, j+2) || (match && rec(i+1, j))
			return mem[Tuple{i, j}]
		}

		if match {
			mem[Tuple{i, j}] = rec(i+1, j+1)
			return mem[Tuple{i, j}]
		}
		mem[Tuple{i, j}] = false
		return false
	}
	return rec(0, 0)
}

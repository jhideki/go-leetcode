package main

func isinterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	type Tuple struct {
		i1 int
		i2 int
	}

	var m = make(map[Tuple]bool)
	var rec func(int, int) bool

	rec = func(i1 int, i2 int) bool {
		if i1 == len(s1) && i2 == len(s2) {
			return true
		}

		if val, exists := m[Tuple{i1, i2}]; exists {
			return val
		}

		if i1 < len(s1) && s3[i1+i2] == s1[i1] && rec(i1+1, i2) {
			return true
		}

		if i2 < len(s2) && s3[i1+i2] == s2[i2] && rec(i1, i2+1) {
			return true
		}

		m[Tuple{i1, i2}] = false

		return false

	}
	return rec(0, 0)
}

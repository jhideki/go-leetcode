package main

func longestIncreasingPath(matrix [][]int) int {

	type Tuple struct {
		x int
		y int
	}

	var m = make(map[Tuple]int)

	var rec func(x int, y int, prev int) int

	rec = func(x int, y int, prev int) int {
		if x >= len(matrix) || y >= len(matrix[0]) || x < 0 || y < 0 {
			return 0
		}

		if val, exists := m[Tuple{x, y}]; exists {
			return val
		}

		if matrix[x][y] <= prev {
			return 0
		}

		prev = matrix[x][y]

		m[Tuple{x, y}] = max(1+rec(x, y+1, prev), 1+rec(x+1, y, prev), 1+rec(x-1, y, prev), 1+rec(x, y-1, prev))
		return m[Tuple{x, y}]
	}

	res := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			res = max(res, rec(i, j, -1))
		}
	}
	return res
}

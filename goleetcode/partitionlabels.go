package main

func partitionLables(s string) []int {
	count := make(map[rune]int)
	res := []int{}
	size, end := 1, 0

	for i, char := range s {
		count[char] = i
	}

	for i, char := range s {
		end = max(end, count[char])
		if i == end {
			res = append(res, size)
			size = 1
		} else {
			size++
		}

	}
	return res
}

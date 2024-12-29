package main

func mergeTriplets(triplets [][]int, target []int) bool {
	x, y, z := false, false, false

	for _, t := range triplets {
		x = x || (t[0] == target[0] && t[1] <= target[1] && t[2] <= target[2])
		y = y || (t[1] == target[1] && t[0] <= target[0] && t[2] <= target[2])
		z = z || (t[2] == target[2] && t[1] <= target[1] && t[0] <= target[0])
		if x && y && z {
			return true
		}
	}
	return false
}

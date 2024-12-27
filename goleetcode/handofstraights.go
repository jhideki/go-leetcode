package main

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	count := map[int]int{}
	for _, num := range hand {
		count[num]++
	}

	for _, val := range hand {
		start := val
		if count[val] == 0 {
			continue
		}
		//Find the min of the hashmap for each iteration
		for count[start-1] > 0 {
			start--
		}
		for start <= val {
			for count[start] > 0 {
				for i := start; i < start+groupSize; i++ {
					if count[i] == 0 {
						return false
					}
					count[i]--
				}
			}
			start++
		}
	}
	return true

}

package main

func maxSubArray(nums []int) int {
	sum := 0
	max := -1 << 31
	for _, value := range nums {
		if sum < 0 && value > sum {
			sum = value
		} else {
			sum += value
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

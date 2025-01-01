package main

func checkValidString(s string) bool {
	leftMin, leftMax := 0, 0

	for _, char := range s {
		if char == '(' {
			leftMin++
			leftMax++
		} else if char == '*' {
			leftMin--
			leftMax++
		} else {
			leftMax--
			leftMin--
		}

		if leftMax < 0 {
			return false
		}
		if leftMin < 0 {
			leftMin = 0
		}
	}
	return leftMin == 0

}

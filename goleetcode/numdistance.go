package main

func minDistance(word1 string, word2 string) int {

	mem := make([][]int, len(word1)+1)

	for i := 0; i < len(word1)+1; i++ {
		mem[i] = make([]int, len(word2)+1)
	}

	for j := 0; j < len(word2)+1; j++ {
		mem[len(word1)][j] = len(word2) - j
	}

	for i := 0; i < len(word1)+1; i++ {
		mem[i][len(word2)] = len(word1) - i
	}

	for i := len(word1) - 1; i >= 0; i-- {
		for j := len(word2) - 1; j >= 0; j-- {
			if word1[i] == word2[j] {
				mem[i][j] = mem[i+1][j+1]
			} else {
				mem[i][j] = min(mem[i+1][j], mem[i][j+1], mem[i+1][j+1])
			}

		}
	}

	return mem[0][0]

}

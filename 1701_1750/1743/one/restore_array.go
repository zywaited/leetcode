package one

func RestoreArray(adjacentPairs [][]int) []int {
	pairs := make(map[int][]int, len(adjacentPairs)+1)
	for _, adjacentPair := range adjacentPairs {
		pairs[adjacentPair[0]] = append(pairs[adjacentPair[0]], adjacentPair[1])
		pairs[adjacentPair[1]] = append(pairs[adjacentPair[1]], adjacentPair[0])
	}
	root := 0
	for num, pair := range pairs {
		if len(pair) == 1 {
			root = num
			break
		}
	}
	num := root
	ans := make([]int, 0, len(adjacentPairs)+1)
	for {
		ans = append(ans, num)
		if len(ans) > len(adjacentPairs) {
			return ans
		}
		for _, pair := range pairs[num] {
			if len(ans) <= 1 || ans[len(ans)-2] != pair {
				num = pair
				break
			}
		}
	}
}

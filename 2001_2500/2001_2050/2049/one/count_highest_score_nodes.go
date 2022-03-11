package one

func CountHighestScoreNodes(parents []int) int {
	children := make([][]int, len(parents))
	for c, p := range parents {
		if p >= 0 {
			children[p] = append(children[p], c)
		}
	}
	ans := 0
	ms := 0
	var dfs func(i int) int
	dfs = func(i int) int {
		var nums [2]int
		for ni, c := range children[i] {
			nums[ni] = dfs(c)
		}
		num := nums[0] + nums[1] + 1
		cm := max(len(parents)-num, 1) * max(nums[0], 1) * max(nums[1], 1)
		if ms < cm {
			ans = 0
			ms = cm
		}
		if ms == cm {
			ans++
		}
		return num
	}
	dfs(0)
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}

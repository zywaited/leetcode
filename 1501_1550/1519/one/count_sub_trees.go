package one

func CountSubTrees(n int, edges [][]int, labels string) []int {
	ans := make([]int, n)
	points := make([][]int, n)
	for _, edge := range edges {
		points[edge[0]] = append(points[edge[0]], edge[1])
		points[edge[1]] = append(points[edge[1]], edge[0])
	}
	nums := [26]int{}
	var dfs func(i int)
	dfs = func(i int) {
		if ans[i] > 0 {
			return
		}
		ans[i] = 1 // 自身也算
		num := nums[labels[i]-'a']
		for _, point := range points[i] {
			dfs(point)
		}
		ans[i] += nums[labels[i]-'a'] - num
		nums[labels[i]-'a']++
	}
	dfs(0)
	return ans
}

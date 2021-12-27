package one

func SumOfDistancesInTree(N int, edges [][]int) []int {
	ans := make([]int, N)
	// 每个节点能正向到达数量
	nn := make(map[int]int, N)

	nm := make([][]int, N)
	for _, edge := range edges {
		nm[edge[0]] = append(nm[edge[0]], edge[1])
		nm[edge[1]] = append(nm[edge[1]], edge[0])
	}

	// 正向计算
	var dfs func(int, int)
	dfs = func(node, prevNode int) {
		nn[node] += 1
		if len(nm[node]) == 0 {
			return
		}
		for _, nextNode := range nm[node] {
			if nextNode != prevNode {
				dfs(nextNode, node)
				nn[node] += nn[nextNode]
				ans[node] += ans[nextNode] + nn[nextNode]
			}
		}
	}
	dfs(0, -1)

	// 方向计算
	var rdfs func(int, int)
	rdfs = func(node, prevNode int) {
		for _, nextNode := range nm[node] {
			if nextNode != prevNode {
				ans[nextNode] += ans[node] - ans[nextNode] - nn[nextNode] + N - nn[nextNode]
				rdfs(nextNode, node)
			}
		}
	}
	rdfs(0, -1)

	return ans
}

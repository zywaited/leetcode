package one

// 这个中文题目翻译难看懂，建议看英文版的
//「离陆地区域最远」要求海洋区域距离它最近的陆地区域的曼哈顿距离是最大的。所以我们需要找一个海洋区域，满足它到陆地的最小距离是最大的。
// 采用bfs找到最近的陆地，然后取最大值
// 这里可以直接从多陆地出发，相比单纯的bfs会好一点

var directs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func MaxDistance(grid [][]int) int {
	type node struct {
		ss int // 陆地的位置
		cs int // 当前位置
	}
	score := 0
	// 其实可以在原数组改，这里只是为了不破坏原数据
	visited := make(map[int]byte)
	queue := make([]*node, 0)
	cnt := 0 // 待找的海洋
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				score = i*101 + j
				queue = append(queue, &node{ss: score, cs: score})
				visited[score] = 1
				continue
			}
			cnt++
		}
	}
	ans := -1
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		for _, direct := range directs {
			ni, nj := n.cs/101+direct[0], n.cs%101+direct[1]
			if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[ni]) {
				continue
			}
			score = ni*101 + nj
			if grid[ni][nj] == 1 || visited[score] == 1 {
				continue
			}
			// 找到一个新海洋
			visited[score] = 1
			ans = max(ans, abs(n.ss/101-ni)+abs(n.ss%101-nj))
			cnt--
			// 所有的都被找到
			if cnt == 0 {
				return ans
			}
			queue = append(queue, &node{ss: n.ss, cs: score})
		}
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

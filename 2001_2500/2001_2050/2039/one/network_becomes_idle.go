package one

func NetworkBecomesIdle(edges [][]int, patience []int) int {
	netEdges := make([][]int, len(patience))
	for _, edge := range edges {
		netEdges[edge[0]] = append(netEdges[edge[0]], edge[1])
		netEdges[edge[1]] = append(netEdges[edge[1]], edge[0])
	}
	type node struct {
		net  int
		time int
	}
	queue := []node{{time: 0}}
	visited := make([]bool, len(patience))
	visited[0] = true
	ans := 0
	for len(queue) > 0 {
		cn := queue[0]
		queue = queue[1:]
		for _, pi := range netEdges[cn.net] {
			if !visited[pi] {
				visited[pi] = true
				queue = append(queue, node{net: pi, time: cn.time + 1})
				c := (cn.time + 1) << 1
				s := c
				if c%patience[pi] == 0 {
					s += c - patience[pi]
				} else {
					s += c - c%patience[pi]
				}
				ans = max(ans, s+1)
			}
		}
	}
	return ans
}

func max(f, s int) int {
	if s <= f {
		return f
	}
	return s
}

package one

func LongestCycle(edges []int) int {
	steps := make([]int, len(edges))
	for i := range steps {
		steps[i] = -1
	}
	visited := make([]int, len(edges))
	ans := -1
	for node := range edges {
		visit := node + 1
		step := 0
		for ; node >= 0 && steps[node] < 0 && visited[node] != visit; node = edges[node] {
			steps[node] = step
			visited[node] = visit
			step++
		}
		if node == -1 || visited[node] != visit {
			continue
		}
		length := step - steps[node]
		if ans < length {
			ans = length
		}
	}
	return ans
}

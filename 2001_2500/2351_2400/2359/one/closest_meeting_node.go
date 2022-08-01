package one

func ClosestMeetingNode(edges []int, node1 int, node2 int) int {
	steps := make([]int, len(edges))
	for i := range steps {
		steps[i] = -1
	}
	step := 0
	node := node1
	for node >= 0 && steps[node] < 0 {
		steps[node] = step
		step++
		node = edges[node]
	}
	ansNode := -1
	ansStep := len(edges) + 1
	step = -1
	node = node2
	visited := make([]bool, len(edges))
	for ; node >= 0 && !visited[node]; node = edges[node] {
		visited[node] = true
		step++
		if steps[node] == -1 {
			continue
		}
		cs := max(step, steps[node])
		if ansStep < cs {
			continue
		}
		if cs < ansStep {
			ansNode = node
			ansStep = cs
			continue
		}
		if ansNode == -1 || node < ansNode {
			ansNode = node
		}
	}
	return ansNode
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}

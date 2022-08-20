package one

func EdgeScore(edges []int) int {
	ans := 0
	scores := make([]int, len(edges))
	for i := 0; i < len(edges); i++ {
		prev := i
		next := edges[prev]
		for next != -1 {
			scores[next] += prev
			if scores[ans] < scores[next] || (scores[ans] == scores[next] && next < ans) {
				ans = next
			}
			edges[prev] = -1
			prev = next
			next = edges[prev]
		}
	}
	return ans
}

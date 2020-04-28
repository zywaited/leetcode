package one

func MinJump(jump []int) int {
	visited := make(map[int]byte, len(jump))
	q := make([][2]int, 0, len(jump))
	q = append(q, [2]int{0, 0})
	visited[0] = 1
	// 该进行的索引
	si := 0
	for len(q) > 0 {
		i, step := q[0][0], q[0][1]
		q = q[1:]
		if i+jump[i] >= len(jump) {
			return step + 1
		}
		// 下一个直接扔进去
		if visited[i+jump[i]] == 0 {
			visited[i+jump[i]] = 1
			q = append(q, [2]int{i + jump[i], step + 1})
		}
		// 左边
		for j := si; j < i; j++ {
			if visited[j] == 0 {
				visited[j] = 1
				q = append(q, [2]int{j, step + 1})
			}
		}
		if i+1 > si {
			si = i + 1
		}
	}
	return -1
}

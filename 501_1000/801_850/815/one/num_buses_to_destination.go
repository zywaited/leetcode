package one

func NumBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	type node struct {
		index int
		step  int
	}
	sm := make(map[int][]int)
	vm := make([]bool, len(routes))
	te := false
	var q []node
	for i := range routes {
		for j := range routes[i] {
			sm[routes[i][j]] = append(sm[routes[i][j]], i)
			if routes[i][j] == source {
				q = append(q, node{index: i, step: 1})
				vm[i] = true
				continue
			}
			if !te && routes[i][j] == target {
				te = true
			}
		}
	}
	if !te {
		return -1
	}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		for _, st := range routes[n.index] {
			if st == target {
				return n.step
			}
			for _, ni := range sm[st] {
				if !vm[ni] {
					vm[ni] = true
					q = append(q, node{index: ni, step: n.step + 1})
				}
			}
		}

	}
	return -1
}

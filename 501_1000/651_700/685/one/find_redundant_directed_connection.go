package one

func FindRedundantDirectedConnection(edges [][]int) []int {
	indexes := make([]int, len(edges)+1)
	for index := range indexes {
		indexes[index] = index
	}

	parent := make([]int, len(indexes))
	copy(parent, indexes)

	var (
		conflict []int
		cycle    []int
	)
	for _, edge := range edges {
		if parent[edge[1]] != edge[1] {
			// 多个父节点
			conflict = edge
			continue
		}
		parent[edge[1]] = edge[0]
		if find(indexes, edge[0]) == find(indexes, edge[1]) {
			cycle = edge
			continue
		}
		union(indexes, edge[0], edge[1])
	}

	if len(conflict) == 0 {
		return cycle
	}
	if len(cycle) == 0 {
		return conflict
	}
	return []int{parent[conflict[1]], conflict[1]}
}

func union(indexes []int, fi, si int) {
	indexes[find(indexes, fi)] = indexes[find(indexes, si)]
}

func find(indexes []int, index int) int {
	for indexes[index] != index {
		indexes[index] = indexes[indexes[index]]
		index = indexes[index]
	}
	return index
}

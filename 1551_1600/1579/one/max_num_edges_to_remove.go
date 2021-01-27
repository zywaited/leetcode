package one

func MaxNumEdgesToRemove(n int, edges [][]int) int {
	aIndexes := make([]int, n+1)
	bIndexes := make([]int, n+1)
	for index := 0; index < n; index++ {
		aIndexes[index] = index
		bIndexes[index] = index
	}
	find := func(indexes []int, index int) int {
		for indexes[index] != index {
			indexes[index] = indexes[indexes[index]]
			index = indexes[index]
		}
		return index
	}
	union := func(indexes []int, firstIndex, secondIndex int) {
		indexes[find(indexes, firstIndex)] = indexes[find(indexes, secondIndex)]
	}
	lined := func(indexes []int, firstIndex, secondIndex int) bool {
		return find(indexes, firstIndex) == find(indexes, secondIndex)
	}
	an := n - 1
	bn := n - 1
	ans := len(edges)
	// 先使用双向边
	for _, edge := range edges {
		if edge[0] == 3 && (!lined(aIndexes, edge[1], edge[2]) || !lined(bIndexes, edge[1], edge[2])) {
			union(aIndexes, edge[1], edge[2])
			union(bIndexes, edge[1], edge[2])
			an--
			bn--
			ans--
		}
	}
	// 考虑是否使用单向边
	for _, edge := range edges {
		switch edge[0] {
		case 1:
			if !lined(aIndexes, edge[1], edge[2]) {
				an--
				ans--
				union(aIndexes, edge[1], edge[2])
			}
		case 2:
			if !lined(bIndexes, edge[1], edge[2]) {
				bn--
				ans--
				union(bIndexes, edge[1], edge[2])
			}
		}
	}
	if an > 0 || bn > 0 {
		return -1
	}
	return ans
}

package one

func FindRedundantConnection(edges [][]int) []int {
	indexes := make([]int, len(edges)+1)
	for index := range indexes {
		indexes[index] = index
	}
	find := func(index int) int {
		for indexes[index] != index {
			indexes[index] = indexes[indexes[index]]
			index = indexes[index]
		}
		return index
	}
	union := func(firstIndex, secondIndex int) {
		indexes[find(firstIndex)] = indexes[find(secondIndex)]
	}
	var conflict []int
	for _, edge := range edges {
		if find(edge[0]) == find(edge[1]) {
			// 冲突了【只有一条边，直接返回】
			conflict = edge
			break
		}
		union(edge[0], edge[1])
	}
	return conflict
}

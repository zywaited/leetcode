package one

func EquationsPossible(equations []string) bool {
	// 每个的上级索引
	indexes := make([]int, 26)
	// 默认上级是自己
	for index := range indexes {
		indexes[index] = index
	}
	// 先合并节点
	for _, equation := range equations {
		if equation[1] == '=' {
			union(indexes, int(equation[0]-'a'), int(equation[3]-'a'))
		}
	}
	// 再判断不等节点是否已经在同一父节点下
	for _, equation := range equations {
		if equation[1] == '!' && find(indexes, int(equation[0]-'a')) == find(indexes, int(equation[3]-'a')) {
			return false
		}
	}
	return true
}

// 合并两个节点(最上层父节点)
func union(indexes []int, fi, si int) {
	indexes[find(indexes, fi)] = indexes[find(indexes, si)]
}

// 查找节点的最上层父节点
func find(indexes []int, index int) int {
	// 如果节点的父节点不是自己，那么就往上查找
	for indexes[index] != index {
		indexes[index] = indexes[indexes[index]] // 这一步作替换后，减少后续查询次数
		index = indexes[index]
	}
	return index
}

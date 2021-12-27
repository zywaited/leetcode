package one

import "sort"

func SmallestStringWithSwaps(s string, pairs [][]int) string {
	// 连通点合并
	indexes := make([]int, len(s))
	for index := range s {
		indexes[index] = index
	}
	find := func(index int) int {
		for indexes[index] != index {
			indexes[index] = indexes[indexes[index]]
			index = indexes[index]
		}
		return indexes[index]
	}
	union := func(firstIndex, secondIndex int) {
		indexes[find(firstIndex)] = indexes[find(secondIndex)]
	}
	for _, pair := range pairs {
		union(pair[0], pair[1])
	}

	// 找出所有的连通节点，对每组节点进行排序填充
	groups := make([][]byte, len(s))
	for index := range s {
		groupIndex := find(index)
		groups[groupIndex] = append(groups[groupIndex], s[index])
	}
	for index := range groups {
		sort.Slice(groups[index], func(i, j int) bool {
			return groups[index][i] < groups[index][j]
		})
	}

	// 组装结果
	ans := make([]byte, len(s))
	for index := range ans {
		groupIndex := find(index)
		ans[index] = groups[groupIndex][0]
		groups[groupIndex] = groups[groupIndex][1:]
	}
	return string(ans)
}

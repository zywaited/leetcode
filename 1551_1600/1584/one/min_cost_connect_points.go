package one

import "sort"

func MinCostConnectPoints(points [][]int) int {
	type node struct {
		i int
		j int
		d int
	}
	nodes := make([]node, 0)
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			nodes = append(nodes, node{i: i, j: j, d: abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1])})
		}
	}
	indexes := make([]int, len(nodes))
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
	for index := range indexes {
		indexes[index] = index
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].d < nodes[j].d
	})
	ans := 0
	for _, n := range nodes {
		if find(n.i) != find(n.j) {
			union(n.i, n.j)
			ans += n.d
		}
	}
	return ans
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

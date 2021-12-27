package one

func MakeConnected(n int, connections [][]int) int {
	indexes := make([]int, n)
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
	line := 0
	over := n - 1
	for _, connection := range connections {
		if find(connection[0]) == find(connection[1]) {
			line++
			continue
		}
		union(connection[0], connection[1])
		over--
	}
	if line < over {
		return -1
	}
	return over
}

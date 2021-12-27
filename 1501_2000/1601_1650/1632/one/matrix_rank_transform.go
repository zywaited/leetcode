package one

import "sort"

func MatrixRankTransform(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	indexes := make([][2]int, 0, m*n)
	for i := range matrix {
		for j := range matrix[i] {
			indexes = append(indexes, [2]int{i, j})
		}
	}
	sort.Slice(indexes, func(i, j int) bool {
		return matrix[indexes[i][0]][indexes[i][1]] < matrix[indexes[j][0]][indexes[j][1]]
	})
	unionIndexes := make([]int, m+n+1)
	unionIndex := 0
	for unionIndex = range unionIndexes {
		unionIndexes[unionIndex] = unionIndex
	}
	find := func(unionIndex int) int {
		for unionIndexes[unionIndex] != unionIndex {
			unionIndexes[unionIndex] = unionIndexes[unionIndexes[unionIndex]]
			unionIndex = unionIndexes[unionIndex]
		}
		return unionIndexes[unionIndex]
	}
	union := func(firstIndex, secondIndex int) {
		unionIndexes[find(firstIndex)] = unionIndexes[find(secondIndex)]
	}
	i := 0
	vm := make([][2]int, m+n)
	mr := make([]int, m)
	nr := make([]int, n)
	for i < len(indexes) {
		j := i
		v := i + 1
		for ; i < len(indexes) && matrix[indexes[i][0]][indexes[i][1]] == matrix[indexes[j][0]][indexes[j][1]]; i++ {
			union(indexes[i][0], m+indexes[i][1])
			unionIndex = find(indexes[i][0])
			if vm[unionIndex][0] < v {
				vm[unionIndex][0] = v
				vm[unionIndex][1] = 0
			}
			if mr[indexes[i][0]] > vm[unionIndex][1] {
				vm[unionIndex][1] = mr[indexes[i][0]]
			}
			if nr[indexes[i][1]] > vm[unionIndex][1] {
				vm[unionIndex][1] = nr[indexes[i][1]]
			}
		}
		r := 0
		for ; j < i; j++ {
			r = vm[find(indexes[j][0])][1] + 1
			matrix[indexes[j][0]][indexes[j][1]] = r
			mr[indexes[j][0]] = r
			nr[indexes[j][1]] = r
		}
	}
	return matrix
}

package one

import "container/list"

var directs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func UpdateMatrix(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	visited := make(map[int]byte, m*n)
	cnt := 0
	queue := &list.List{}
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				sc := i*10001 + j
				visited[sc] = 1
				queue.PushBack([2]int{sc, 0})
				continue
			}
			cnt++
		}
	}
	// 无1不处理
	if cnt == 0 {
		return matrix
	}
	for queue.Len() > 0 {
		cn := queue.Front().Value.([2]int)
		queue.Remove(queue.Front())
		for _, direct := range directs {
			i, j := cn[0]/10001+direct[0], cn[0]%10001+direct[1]
			if i < 0 || i >= m || j < 0 || j >= n || matrix[i][j] == 0 {
				continue
			}
			sc := i*10001 + j
			if visited[sc] == 1 {
				continue
			}
			// 找到最近的1
			visited[sc] = 1
			matrix[i][j] = cn[1] + 1
			cnt--
			if cnt == 0 {
				return matrix
			}
			queue.PushBack([2]int{sc, matrix[i][j]})
		}
	}
	return matrix
}

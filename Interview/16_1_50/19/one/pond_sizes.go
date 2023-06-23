package one

import "sort"

var directs = [][]int{
	{0, -1},
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
}

func PondSizes(land [][]int) []int {
	sizes := make([]int, 0)
	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[i]); j++ {
			if land[i][j] == 0 {
				sizes = append(sizes, pondSize(land, i, j))
			}
		}
	}
	sort.Ints(sizes)
	return sizes
}

func pondSize(land [][]int, i, j int) int {
	size := 0
	land[i][j] = 1
	queue := [][]int{{i, j}}
	for len(queue) > 0 {
		size++
		point := queue[0]
		queue = queue[1:]
		ci, cj := point[0], point[1]
		for _, direct := range directs {
			ni, nj := ci+direct[0], cj+direct[1]
			if ni >= 0 && nj >= 0 && ni < len(land) && nj < len(land[ni]) && land[ni][nj] == 0 {
				land[ni][nj] = 1
				queue = append(queue, []int{ni, nj})
			}
		}
	}
	return size
}

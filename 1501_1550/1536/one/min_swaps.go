package one

import "container/list"

func MinSwaps(grid [][]int) int {
	gl := list.New()
	for i := range grid {
		zn := 0
		for j := len(grid[i]) - 1; j >= 0 && grid[i][j] == 0; j-- {
			zn++
		}
		gl.PushBack(zn)
	}
	step := 0
	fn := 0
	for gn := len(grid) - 1; gn > 0; gn-- {
		cn := gl.Front()
		fn = 0
		for ; cn != nil && cn.Value.(int) < gn; cn = cn.Next() {
			fn++
		}
		if cn == nil {
			return -1
		}
		gl.Remove(cn)
		step += fn
	}
	return step
}

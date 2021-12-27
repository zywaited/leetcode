package one

func ColorBorder(grid [][]int, r int, c int, color int) [][]int {
	// 重新生成返回值(不然更改其中的数值会影响程序判断逻辑)
	gds := make([][]int, len(grid))
	for index, gd := range grid {
		gds[index] = make([]int, len(gd))
		copy(gds[index], gd)
	}
	rm := len(gds) - 1
	cm := len(gds[0]) - 1
	// 记录原始颜色和已经记录的颜色
	bc := grid[r][c]
	colored := make(map[[2]int]byte)
	var cb func(int, int)
	cb = func(cr, cc int) {
		// 是否使用过了
		key := [2]int{cr, cc}
		if colored[key] == 1 {
			return
		}
		colored[key] = 1
		// 4个方向判断
		// 是否处于边界
		border := uint8(0)
		// n
		if cr > 0 && grid[cr-1][cc] == bc {
			cb(cr-1, cc)
			border++
		}
		// l
		if cc > 0 && grid[cr][cc-1] == bc {
			cb(cr, cc-1)
			border++
		}
		// w
		if cr < rm && grid[cr+1][cc] == bc {
			cb(cr+1, cc)
			border++
		}
		// r
		if cc < cm && grid[cr][cc+1] == bc {
			cb(cr, cc+1)
			border++
		}
		if border < 4 {
			// 此处<4判断已经包含了最外层
			gds[cr][cc] = color
		}
	}
	cb(r, c)
	return gds
}

package one

func CanMeasureWater(x int, y int, z int) bool {
	// 已经尝试过的组合
	used := make(map[[2]int]byte)
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{0, 0})
	for len(queue) > 0 {
		cn := queue[0]
		queue = queue[1:]
		if cn[0] == z || cn[1] == z || cn[0]+cn[1] == z {
			return true
		}
		if used[cn] == 1 {
			continue
		}
		used[cn] = 1
		// 清空x
		queue = append(queue, [2]int{0, cn[1]})
		// 清空y
		queue = append(queue, [2]int{cn[0], 0})
		// 灌满x
		queue = append(queue, [2]int{x, cn[1]})
		// 灌满y
		queue = append(queue, [2]int{cn[0], y})
		// x往y灌水
		queue = append(queue, [2]int{cn[0] - min(y-cn[1], cn[0]), cn[1] + min(y-cn[1], cn[0])})
		// y往x灌水
		queue = append(queue, [2]int{cn[0] + min(x-cn[0], cn[1]), cn[1] - min(x-cn[0], cn[1])})
	}
	return false
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}

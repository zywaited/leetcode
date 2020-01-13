package one

func Robot(command string, obstacles [][]int, x int, y int) bool {
	// 起点也算在里面
	mp := map[int64]byte{0: 1}
	xs, ys := 0, 0
	for i := range command {
		switch command[i] {
		case 'U':
			ys++
		case 'R':
			xs++
		}
		mp[int64(xs*(1e9+1)+ys)] = 1
	}
	// 判断终点是否在集合中
	c := 0
	c = min(x/xs, y/ys)
	if mp[int64((x-(xs*c))*(1e9+1)+(y-(ys*c)))] == 0 {
		return false
	}
	// 判断障碍物是否在集合中
	for _, pointer := range obstacles {
		// 如果超出了终点就忽略
		if pointer[0] > x || pointer[1] > y {
			continue
		}
		c = min(pointer[0]/xs, pointer[1]/ys)
		if mp[int64((pointer[0]-(xs*c))*(1e9+1)+(pointer[1]-(ys*c)))] == 1 {
			return false
		}
	}
	return true
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}

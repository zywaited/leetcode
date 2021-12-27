package one

// 同一条直线斜率是一样的
func MaxPoints(points [][]int) int {
	key := [2]int{} // 斜率
	max := 0        // 最大的点数
	// 以下逻辑有很多重复计算过程, 后续可以考虑优化该部分
	for i := 0; i < len(points); i++ {
		same := 0                     // 直线有相同的点存在
		size := 1                     // 竖线上的点数
		cm := 0                       // 当前最大值
		mc := make(map[[2]int]int, 0) // 当前各个数量
		for j := 0; j < len(points); j++ {
			if i != j {
				// x相同
				if points[i][0] == points[j][0] {
					// y相同
					if points[i][1] == points[j][1] {
						same++
					}
					size++
					continue
				}
				// 获取斜率
				// [2]int代表斜率, 为了保证无法除尽的情况(精度可能导致结果不一致)
				// 通过gcd找出最大公约数后[2]int就是不能约分后的最终斜率
				dx := points[j][0] - points[i][0]
				dy := points[j][1] - points[i][1]
				gcd := gcd(dx, dy)
				if gcd != 0 {
					dx = dx / gcd
					dy = dy / gcd
				}
				key[0] = dy
				key[1] = dx
				// 直线点数加1
				mc[key]++
				if mc[key] > cm {
					cm = mc[key]
				}
			}
		}
		// 最终点数
		cm += same + 1
		if size > cm {
			cm = size
		}
		// 获取最大值
		if cm > max {
			max = cm
		}
	}
	return max
}

// 欧几里得算法找最大公约数
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

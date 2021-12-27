package two

import "math"

// 按照对角线计算面积
func MinAreaRect(points [][]int) int {
	// 存储每个点(判断对角线就是其余两个点都存在)
	lm := make(map[int]byte)
	for _, point := range points {
		lm[point[0]*40001+point[1]] = 1
	}
	area := math.MaxInt32
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			// 两个点必须要不在同一个x轴或者y轴
			if points[i][0] != points[j][0] && points[i][1] != points[j][1] {
				// 其余两个点存在
				if lm[points[i][0]*40001+points[j][1]] == 1 && lm[points[j][0]*40001+points[i][1]] == 1 {
					area = min(area, abs((points[i][0]-points[j][0])*(points[i][1]-points[j][1])))
				}
			}
		}
	}
	if area == math.MaxInt32 {
		return 0
	}
	return area
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

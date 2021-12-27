package one

import (
	"math"
	"sort"
)

// 按照x轴分类(防止重复计算)，然后一个一个点找最近的矩形
func MinAreaRect(points [][]int) int {
	sort.Sort(ps(points))
	// 同个X轴的两个点(距离最近有数据的x轴)
	// 唯一值确认: y1 * 40001 + y2
	lm := make(map[int]int)
	// x对应y数据集合
	xym := make([][][]int, 0)
	x := -1
	i := -1
	area := math.MaxInt32
	for _, point := range points {
		if x != point[0] {
			xym = append(xym, [][]int{point})
			i = len(xym) - 1
			x = point[0]
			continue
		}
		xym[i] = append(xym[i], point)
	}
	for _, xys := range xym {
		for i := range xys {
			for j := i + 1; j < len(xys); j++ {
				v := xys[i][1]*40001 + xys[j][1]
				if x, ok := lm[v]; ok {
					area = min(area, (xys[j][0]-x)*(xys[j][1]-xys[i][1]))
				}
				lm[v] = xys[j][0]
			}
		}
	}
	if area == math.MaxInt32 {
		return 0
	}
	return area
}

type ps [][]int

func (ps ps) Len() int {
	return len(ps)
}

func (ps ps) Less(i, j int) bool {
	// 先判断x轴
	if ps[i][0] > ps[j][0] {
		return false
	}
	if ps[i][0] < ps[j][0] {
		return true
	}
	// 判断y轴
	return ps[i][1] < ps[j][1]
}

func (ps ps) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}

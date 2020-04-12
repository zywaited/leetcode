package one

func Intersection(start1 []int, end1 []int, start2 []int, end2 []int) []float64 {
	// 先优先处理垂直线
	if start1[0] == end1[0] || start2[0] == end2[0] {
		// 平行
		if start1[0] == end1[0] && start2[0] == end2[0] {
			if start1[0] != start2[0] {
				return nil
			}
			if min(start1[1], end1[1]) > max(start2[1], end2[1]) || max(start1[1], end1[1]) < min(start2[1], end2[1]) {
				return nil
			}
			if min(start1[1], end1[1]) >= min(start2[1], end2[1]) {
				return []float64{float64(start1[0]), float64(min(start1[1], end1[1]))}
			}
			return []float64{float64(start1[0]), float64(min(start2[1], end2[1]))}
		}
		// 交点(y)
		// (end2[1]-y)/(end2[0]-start1[0]) = (end2[1]-start2[1])/(end2[0]-start2[0])
		y := float64(end2[1]) - float64((end2[1]-start2[1])*(end2[0]-start1[0]))/float64(end2[0]-start2[0])
		if diff(float64(start1[0]), y, start1, end1) || diff(float64(start1[0]), y, start2, end2) {
			return nil
		}
		return []float64{float64(start1[0]), y}
	}
	// 斜率
	g1 := gcd(abs(end1[1]-start1[1]), abs(end1[0]-start1[0]))
	r1 := [3]int{abs(end1[1]-start1[1]) / g1, abs(end1[0]-start1[0]) / g1, 1}
	if (end1[1]-start1[1])*(end1[0]-start1[0]) < 0 {
		r1[2] = -1
	}
	g2 := gcd(abs(end2[1]-start2[1]), abs(end2[0]-start2[0]))
	r2 := [3]int{abs(end2[1]-start2[1]) / g2, abs(end2[0]-start2[0]) / g2, 1}
	if (end2[1]-start2[1])*(end2[0]-start2[0]) < 0 {
		r2[2] = -1
	}
	// 平行
	if r1 == r2 {
		// 如果line1随便一点都在斜率内就重叠
		g3 := gcd(abs(end2[1]-end1[1]), abs(end2[0]-end1[0]))
		if g3 > 0 {
			r3 := [3]int{abs(end2[1]-end1[1]) / g3, abs(end2[0]-end1[0]) / g3, 1}
			if (end2[1]-end1[1])*(end2[0]-end1[0]) < 0 {
				r3[2] = -1
			}
			if r1 != r3 {
				return nil
			}
		}
		if min(start1[0], end1[0]) > max(start2[0], end2[0]) || max(start1[0], end1[0]) < min(start2[0], end2[0]) {
			return nil
		}
		if min(start1[0], end1[0]) >= min(start2[0], end2[0]) {
			if start1[0] < end1[0] {
				return []float64{float64(start1[0]), float64(start1[1])}
			}
			return []float64{float64(end1[0]), float64(end1[1])}
		}
		if start2[0] < end2[0] {
			return []float64{float64(start2[0]), float64(start2[1])}
		}
		return []float64{float64(end2[0]), float64(end2[1])}
	}
	// x, y
	// (end2[1]-y)/(end2[0]-x) = (end2[1]-start2[1])/(end2[0]-start2[0])
	// (end1[1]-y)/(end1[0]-x) = (end1[1]-start1[1])/(end1[0]-start1[0])
	// (end2[1]-end1[1]+r1*(end1[0]-x)) = r2(end2[0]-x)
	// x = (end2[1]-end1[1]+r1*end1[0]-r2*end2[0]) / (r1-r2)
	x := (float64(end2[1]-end1[1]) + float64(r1[0]*end1[0]*r1[2])/float64(r1[1]) - float64(r2[0]*end2[0]*r2[2])/float64(r2[1])) / (float64(r1[0]*r2[1]*r1[2]-r1[1]*r2[0]*r2[2]) / float64(r1[1]*r2[1]))
	y := float64(end2[1]) - (float64(end2[0])-x)*float64(r2[0]*r2[2])/float64(r2[1])
	if diff(x, y, start1, end1) || diff(x, y, start2, end2) {
		return nil
	}
	return []float64{x, y}
}

func diff(x, y float64, s, e []int) bool {
	return x < float64(min(s[0], e[0])) ||
		x > float64(max(s[0], e[0])) ||
		y < float64(min(s[1], e[1])) ||
		y > float64(max(s[1], e[1]))
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}

func gcd(f, s int) int {
	for s > 0 {
		f, s = s, f%s
	}
	return f
}

func abs(f int) int {
	if f < 0 {
		f = -f
	}
	return f
}

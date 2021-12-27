package one

func SearchMatrix(matrix [][]int, target int) bool {
	// 以x=y为中心就行右下搜索
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	rs := func(r, s, e int) bool {
		for s <= e {
			m := s + (e-s)>>1
			if matrix[r][m] == target {
				return true
			}
			if matrix[r][m] < target {
				s = m + 1
				continue
			}
			e = m - 1
		}
		return false
	}
	cs := func(c, s, e int) bool {
		for s <= e {
			m := s + (e-s)>>1
			if matrix[m][c] == target {
				return true
			}
			if matrix[m][c] < target {
				s = m + 1
				continue
			}
			e = m - 1
		}
		return false
	}
	x := 0
	for ; x < m && x < n; x++ {
		if matrix[x][x] == target {
			return true
		}
		if matrix[x][x] > target {
			break
		}
		if rs(x, x+1, n-1) || cs(x, x+1, m-1) {
			return true
		}
	}
	return false
}

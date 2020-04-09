package one

func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	s := 0
	e := m*n - 1
	for s <= e {
		mid := s + (e-s)>>1
		i, j := mid/n, mid%n
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			s = mid + 1
			// 特殊优化当前列
			if matrix[i][n-1] == target {
				return true
			}
			if matrix[i][n-1] < target {
				s = mid + n - j
			} else {
				e = min(e, mid+n-j-2)
			}
			continue
		}
		e = mid - 1
		// 特殊优化当前列
		if matrix[i][0] == target {
			return true
		}
		if matrix[i][0] < target {
			s = max(s, mid-j+1)
		} else {
			e = mid - j - 1
		}
	}
	return false
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}

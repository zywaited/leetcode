package one

func KthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	l, r := matrix[0][0], matrix[n-1][n-1]
	check := func(m int) bool {
		i, j := n-1, 0
		num := 0
		for i >= 0 && j < n {
			if matrix[i][j] <= m {
				num += i + 1
				j++
				continue
			}
			i--
		}
		return num >= k
	}
	for l <= r {
		m := l + (r-l)>>1
		if check(m) {
			r = m - 1
			continue
		}
		l = m + 1
	}
	return l
}

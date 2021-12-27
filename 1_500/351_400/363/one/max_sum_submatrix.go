package one

import "math"

func MaxSumSubmatrix(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	prev := [101]int{}
	current := [101]int{}
	i, j := 1, 1
	sum := 0
	ans := math.MinInt32
	for i <= m && j <= n {
		prev = [101]int{}
		for ci := i; ci <= m; ci++ {
			sum = 0
			for cj := j; cj <= n; cj++ {
				sum += matrix[ci-1][cj-1]
				current[cj] = prev[cj] + sum
				if current[cj] > k {
					continue
				}
				ans = max(ans, current[cj])
				if ans == k {
					return ans
				}
			}
			prev, current = current, prev
		}
		j++
		if j > n {
			i++
			j = 1
			continue
		}
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}

package one

func MaxStudents(seats [][]byte) int {
	// 数组+1为了不判断0
	m, n := len(seats), len(seats[0])
	dp := make([][]int, m+1)
	counts := make([]int, 1<<uint(n))
	bs := make([]int, m+1)
	// 每种情况的人员数量(也就是1的个数)
	for i := range counts {
		counts[i] = counts[i>>1] + i&1 // 1的个数 = 去除最后一位1值的1个数+最后一位是不是1
	}
	// 坏座位
	dp[0] = make([]int, 1<<uint(n)) // 初始化第一列
	for i := 1; i <= m; i++ {
		dp[i] = make([]int, 1<<uint(n))
		for j := 0; j < n; j++ {
			if seats[i-1][j] == '#' {
				bs[i] |= 1 << uint(j)
			}
		}
	}
	for i := 1; i <= m; i++ {
		for j := 0; j < 1<<uint(n); j++ {
			// 判断是否可以安排人(不是坏座位并且左右没有人)
			if j&bs[i] == 0 && j&(j<<1) == 0 && j&(j>>1) == 0 {
				for k := 0; k < 1<<uint(n); k++ {
					// 左右前方不能有人
					if k&(j>>1) == 0 && k&(j<<1) == 0 {
						dp[i][j] = max(dp[i][j], dp[i-1][k]+counts[j]) // 上一排每种情况的数量+当前排当前情况数量
					}
				}
			}
		}
	}
	// 最后一排最大值
	ans := 0
	for j := 0; j < 1<<uint(n); j++ {
		ans = max(ans, dp[m][j])
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}

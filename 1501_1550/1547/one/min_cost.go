package one

import (
	"math"
	"sort"
)

func MinCost(n int, cuts []int) int {
	lc := len(cuts)
	// 加入两端，不加入两端的话计算长度都要做处理
	cuts = append(cuts, 0)
	cuts = append(cuts, n)
	sort.Ints(cuts)
	// dp[i][j] 代表切割i-j区间的最小成本
	dp := make([][]int, len(cuts))
	for i := len(cuts) - 1; i > 0; i-- {
		dp[i] = make([]int, len(cuts))
		if i > lc {
			continue
		}
		for j := i; j <= lc; j++ {
			dp[i][j] = math.MaxInt32
			if i == j {
				dp[i][j] = 0
			}
			for k := i; k <= j; k++ {
				// k是第一次切分在最后体现
				dp[i][j] = min(dp[i][j], dp[i][k-1]+dp[k+1][j])
			}
			// 这里加上长度代表第一次切分
			dp[i][j] += cuts[j+1] - cuts[i-1]
		}
	}
	return dp[1][lc]
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}

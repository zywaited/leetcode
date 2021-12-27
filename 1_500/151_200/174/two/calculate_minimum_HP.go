package two

import "math"

// 对解法1进行空间压缩
// i只依赖于i+1
// 那么临时保留i+1即可
func CalculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 0
	}
	m, n := len(dungeon), len(dungeon[0])
	dp := make([]int, n)
	ni := 0 // [i+1][j]
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			// 如果是终点
			if i+1 == m && j+1 == n {
				if dungeon[i][j] >= 1 {
					dp[j] = 1
				} else {
					dp[j] = 1 - dungeon[i][j]
				}
				continue
			}
			// 现在的dp[j]就是dp[i+1][j]
			ni = dp[j]
			dp[j] = math.MaxInt32
			if i+1 < m {
				// 下方
				if dungeon[i][j] >= ni {
					dp[j] = 1
					continue // 当然这句不要也可以，反正1最小
				} else {
					dp[j] = min(dp[j], ni-dungeon[i][j])
				}
			}
			if j+1 < n {
				// 右方
				if dungeon[i][j] >= dp[j+1] {
					dp[j] = 1
				} else {
					dp[j] = min(dp[j], dp[j+1]-dungeon[i][j])
				}
			}
		}
	}
	return dp[0]
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}

package one

import "math"

// 第一想法肯定是动态规划
// dp[i][j] 代表到达该点所需的最小健康点数
// 这里需要思考下方案:
// 1 左上角--右上角
// 2 右上角--左上角

// 比如
// [-1,-1]
// [-2, 0]
// [4 ,-4]
//
// 第一种方案
// 到达dp[2][0] = 4，dp[1][1] = 3，但是到最终dp[2][1]要选择dp[2][0]，此时是跟所需健康点和剩下的健康点两个数据有关，情况比搞得很复杂
//
// 第二种方案
// dp[2][1] = 5
// dp[2][0] = 1 // 只有剩的健康点就能到到达
// dp[1][1] = 5 //
// 这时的所需的就是只跟能到达下一格的最小健康点有关
// dungeon[i][j] >= dp[i][j+1]; dp[i][j] = 1
// dungeon[i][j] < dp[i][j+1]; dp[i][j]= dp[i][j+1]-dungeon[i][j]
// dp[i+1][j]与之类似，然后取个最小值即可
func CalculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 0
	}
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m)
	for i := m - 1; i >= 0; i-- {
		dp[i] = make([]int, n)
		for j := n - 1; j >= 0; j-- {
			// 如果是终点
			if i+1 == m && j+1 == n {
				if dungeon[i][j] >= 1 {
					dp[i][j] = 1
				} else {
					dp[i][j] = 1 - dungeon[i][j]
				}
				continue
			}
			dp[i][j] = math.MaxInt32
			if i+1 < m {
				// 下方
				if dungeon[i][j] >= dp[i+1][j] {
					dp[i][j] = 1
					continue // 当然这句不要也可以，反正1最小
				} else {
					dp[i][j] = min(dp[i][j], dp[i+1][j]-dungeon[i][j])
				}
			}
			if j+1 < n {
				// 右方
				if dungeon[i][j] >= dp[i][j+1] {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(dp[i][j], dp[i][j+1]-dungeon[i][j])
				}
			}
		}
	}
	return dp[0][0]
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}

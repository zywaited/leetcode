package one

import "sort"

func PileBox(box [][]int) int {
	// 先排序选择最优
	sort.Slice(box, func(i, j int) bool {
		if box[i][2] != box[j][2] {
			return box[i][2] > box[j][2]
		}
		if box[i][1] != box[j][1] {
			return box[i][1] > box[j][1]
		}
		return box[i][0] > box[j][0]
	})
	// i代表第i个箱子，dp[i]代表选择第i个箱子的最大高度
	dp := make([]int, len(box))
	dp[0] = box[0][2]
	ans := box[0][2]
	for i := 1; i < len(box); i++ {
		// 找出可以放置箱子的最大值
		max := 0
		for j := i - 1; j >= 0; j-- {
			if box[j][0] > box[i][0] && box[j][1] > box[i][1] && box[j][2] > box[i][2] && dp[j] > max {
				max = dp[j]
			}
		}
		dp[i] = max + box[i][2]
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

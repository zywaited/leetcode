package one

func FindLatestStep(arr []int, m int) int {
	// dp[i] i对应的左右长度(包含自身)
	dp := make([][2]int, len(arr)+2)
	// 索引是否满足(这里只处理最左)
	dm := make([]bool, len(arr)+2)
	mn := 0
	ans := -1
	for index, num := range arr {
		dp[num][0] = 1
		dp[num][1] = 1

		if dp[num-1][0] > 0 {
			dp[num][0] += dp[num-1][0]
		}
		if dp[num+1][1] > 0 {
			dp[num][1] += dp[num+1][1]
		}
		if dp[num-1][1] > 0 {
			// 更新最左数据
			dp[num-dp[num-1][0]][1] += dp[num][1]
		}
		if dp[num+1][0] > 0 {
			// 更新最右数据
			dp[num+dp[num+1][1]][0] += dp[num][0]
		}
		if dp[num][0]+dp[num][1]-1 == m {
			ans = index + 1
			dm[num-dp[num-1][0]] = true
			mn++
			continue
		}

		if dm[num-dp[num-1][0]] {
			dm[num-dp[num-1][0]] = false
			mn--
		}
		// 右边已经最左数据
		if dm[num+1] {
			dm[num+1] = false
			mn--
		}

		if mn > 0 {
			ans = index + 1
		}
	}
	return ans
}

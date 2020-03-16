package one

func GetKthMagicNumber(k int) int {
	// 1这个数比较特殊
	dp := make([]int, k+1)
	dp[1] = 1
	i3, i5, i7 := 1, 1, 1 // 上一次使用3个数的索引
	for num := 2; num <= k; num++ {
		// 从当前取出最小的那一个数
		// 如果出现乘积相等，那么都往后推一位防止重复数据
		dp[num] = min(dp[i3]*3, min(dp[i5]*5, dp[i7]*7))
		if dp[num] == dp[i3]*3 {
			i3++
		}
		if dp[num] == dp[i5]*5 {
			i5++
		}
		if dp[num] == dp[i7]*7 {
			i7++
		}
	}
	return dp[k]
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}

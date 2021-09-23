package one

// dp[i] 选择合并前i堆石子的最大差值
func StoneGameVIII(stones []int) int {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}
	ans := sum
	for i := len(stones) - 2; i > 0; i-- {
		sum -= stones[i+1]
		ans = max(ans, sum-ans)
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}

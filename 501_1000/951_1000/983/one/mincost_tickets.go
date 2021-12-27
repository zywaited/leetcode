package one

func MincostTickets(days []int, costs []int) int {
	maxDay := days[len(days)-1]
	// 需要计算的天
	de := make(map[int]byte, len(days))
	for _, day := range days {
		de[day] = 1
	}
	dp := make([]int, maxDay+1)
	for day := 1; day <= maxDay; day++ {
		if de[day] == 0 {
			// 与前一天的花费一致(这天不旅行)
			dp[day] = dp[day-1]
			continue
		}
		// 当天花费1天
		dp[day] = dp[day-1] + costs[0]
		// 当天花费7天
		if day-7 >= 0 {
			dp[day] = min(dp[day], dp[day-7]+costs[1])
		} else {
			dp[day] = min(dp[day], costs[1])
		}
		// 当天花费30天
		if day-30 >= 0 {
			dp[day] = min(dp[day], dp[day-30]+costs[2])
		} else {
			dp[day] = min(dp[day], costs[2])
		}
	}
	return dp[maxDay]
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}

package one

func MinRefuelStops(target int, startFuel int, stations [][]int) int {
	// dp[i][j] 到达i处加油j次数所剩的最大燃油量
	dp := make([]int, len(stations)+1)
	dp[0] = startFuel
	pm := startFuel
	tm := 0
	dm := 0
	for i := 1; i <= len(stations); i++ {
		dm = stations[i-1][0] - tm
		if pm < dm {
			return -1
		}
		tm = stations[i-1][0]
		pm = 0
		for j := i; j >= 0; j-- {
			dp[j] = dp[j] - dm
			if dp[j] < 0 {
				dp[j] = -1
			}
			if j > 0 {
				if dp[j-1] < dm {
					continue
				}
				dp[j] = max(dp[j], dp[j-1]-dm+stations[i-1][1])
			}
			pm = max(pm, dp[j])
		}
	}
	dm = target - tm
	for i, f := range dp {
		if f >= dm {
			return i
		}
	}
	return -1
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}

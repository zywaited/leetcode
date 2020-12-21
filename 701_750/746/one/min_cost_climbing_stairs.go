package one

func MinCostClimbingStairs(cost []int) int {
	return min(minCost(cost), minCost(cost[1:]))
}

func minCost(cost []int) int {
	f, s := 0, 0
	t := 0
	for i := 1; i <= len(cost); i++ {
		if i == 1 {
			f = cost[i-1]
			continue
		}
		t = f
		f = min(f+cost[i-1], s+cost[i-2])
		s = t
	}
	return f
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}

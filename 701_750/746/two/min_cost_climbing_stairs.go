package two

func MinCostClimbingStairs(cost []int) int {
	// 简化流程
	// 起点可以为0或者1，只要设置f, s都为0就可以了
	// 可以由第二个节点来选择从哪里开始
	f, s := 0, 0
	t := 0
	for i := 2; i <= len(cost); i++ {
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

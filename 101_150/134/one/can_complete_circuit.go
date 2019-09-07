package one

// 可参考53题
func CanCompleteCircuit(gas []int, cost []int) int {
	// 贪心
	ts := 0 // 总剩余的油
	cs := 0 // 当前剩余的油
	g, s := 0, 0
	for i := 0; i < len(gas); i++ {
		s = gas[i] - cost[i]
		ts += s
		if cs > 0 {
			cs += s
			continue
		}
		cs = s // 下一站为起点
		g = i
	}
	if ts >= 0 && cs >= 0 {
		// 最后有剩余的油并且当前能到下一站
		return g
	}
	return -1
}

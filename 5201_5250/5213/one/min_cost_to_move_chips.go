package one

func MinCostToMoveChips(chips []int) int {
	// 由题目可以知道
	// 所有的位置的筹码都能通用偶数次(代价0)移动到0或者1的位置上
	// 最后比较0,1位置上的数量就行
	f, s := 0, 0
	for _, chip := range chips {
		if chip%2 == 0 {
			s++
			continue
		}
		f++
	}
	if f < s {
		return f
	}
	return s
}

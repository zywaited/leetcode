package one

func MinOperations(n int) int {
	step := 0
	// n就是中心节点
	// 获取需要处理的节点索引
	index := n >> 1
	for cn := index + 1; cn < n; cn++ {
		step += cn<<1 + 1 - n
	}
	if index<<1+1 != n {
		step++
	}
	return step
}

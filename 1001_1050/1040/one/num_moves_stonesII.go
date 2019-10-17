package one

import "sort"

func NumMovesStonesII(stones []int) []int {
	sort.Ints(stones)
	rs := make([]int, 2)
	// 最大值
	// 与移动石子直到连续这道题一样算出长度，因为需要移动端口石子
	// 所以需要减去移动最左边或者最右边的端口石子的最小跳过数
	sn := len(stones)
	// 等价，只是max看着会舒服点
	//rs[1] = stones[sn-1] - stones[0] - sn + 1 - min(stones[sn-1]-stones[sn-2]-1, stones[1]-stones[0]-1)
	rs[1] = max(stones[sn-2]-stones[0], stones[sn-1]-stones[1]) - sn + 2
	rs[0] = rs[1]
	ms := 0 // 最小次数
	i, j := 0, 1
	for ; i < sn; i++ {
		// 找到一个位置区间可以放下所有的石子
		// 这里用<=sn是因为==sn可能是刚好不用移动的情况，这时就是需要判断是否后面还有石子
		for j < sn && stones[j]-stones[i]+1 <= sn {
			j++
		}
		ms = sn - (j - i)
		// 但是给出的实例中[6,5,4,3,10]不能通过这种方法移动
		// 这种情况最小值就是2
		// 特殊处理
		if ms == 1 && stones[j-1]-stones[i]+1 == sn-1 {
			ms = 2
		}
		rs[0] = min(rs[0], ms)
	}
	return rs
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}

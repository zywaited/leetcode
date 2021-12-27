package one

type StockSpanner struct {
	stack [][2]int // 0: 索引 1: 值
	cl    int      // 当前长度
	ll    int      // 总长度
}

func Constructor() StockSpanner {
	// 最大值为1e5
	return StockSpanner{stack: [][2]int{{-1, 1e5 + 1}}}
}

func (ss *StockSpanner) Next(price int) int {
	// 找到前一个大的
	pl := ss.cl
	for ss.stack[pl][1] <= price {
		pl--
	}
	if pl < ss.cl {
		ss.cl = pl
		ss.stack = ss.stack[:ss.cl+1] // 第一个要保留
	}
	ss.cl++
	ss.stack = append(ss.stack, [2]int{ss.ll, price})
	ss.ll++
	return ss.ll - ss.stack[pl][0] - 1
}

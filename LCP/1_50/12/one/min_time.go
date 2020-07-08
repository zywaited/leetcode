package one

func MinTime(time []int, m int) int {
	l := 0
	r := 0
	for _, nt := range time {
		r += nt
	}
	for l <= r {
		t := l + (r-l)>>1
		if valid(time, m, t) {
			r = t - 1
			continue
		}
		l = t + 1
	}
	return l
}

func valid(time []int, m, t int) bool {
	// 在一个区间内去除最大值(场外求助), 不比t大则可以放在一天内
	// 最后判断所需天数是否不大于m即可
	nd := 1
	maxTime := 0   // 当前区间最大值
	costTime := 0  // 当前要花费的时间
	totalTime := 0 // 当前区间的总值
	for _, nt := range time {
		costTime = min(maxTime, nt) // 去除最大时间值
		if costTime+totalTime > t {
			nd++ // 所需天数+1
			maxTime = nt
			totalTime = 0
			continue
		}
		totalTime += costTime
		maxTime = max(maxTime, nt)
	}
	return nd <= m
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}

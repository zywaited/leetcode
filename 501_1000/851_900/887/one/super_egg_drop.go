package one

func SuperEggDrop(K int, N int) int {
	// K个鸡蛋测试N层楼所需的最大移动次数(因为F是所有的可能性，所以这里要算最大的才能包含所有)
	dp := make([][]int, K+1)
	var fc func(int, int) int
	fc = func(tk int, tn int) int {
		if tn <= 0 {
			return 0
		}
		if tk == 1 {
			return tn
		}
		if dp[tk] == nil {
			dp[tk] = make([]int, N+1)
		}
		if dp[tk][tn] > 0 {
			return dp[tk][tn]
		}
		s := 1
		e := tn
		// s==e计算无意义了，肯定是相等的
		for s+1 < e {
			m := (s + e) >> 1
			// 两种情况，鸡蛋没碎和碎了
			a1 := fc(tk, tn-m)
			a2 := fc(tk-1, m-1)
			if a1 == a2 {
				// 结果一样
				s = m
				e = m
				continue
			}
			// 这里不用m-1&m+1也是保证s<e (s==e重复计算也是可以的)
			if a1 < a2 {
				// 最大值在左边
				e = m
				continue
			}
			s = m
		}
		// 比较s与e为分割点的次数
		ans := 1 + min(max(fc(tk, tn-s), fc(tk-1, s-1)), max(fc(tk, tn-e), fc(tk-1, e-1)))
		dp[tk][tn] = ans
		return ans
	}
	return fc(K, N)
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

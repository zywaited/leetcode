package one

func LastRemaining(n int) int {
	k := 1
	l := 1
	r := n
	s := 1
	for cn := n; cn > 1; cn >>= 1 {
		if k&1 > 0 { // 从左到右
			l += s
			if cn&1 > 0 { // 奇数最后一位被去除
				r -= s
			}
		} else { // 从右到左
			r -= s
			if cn&1 > 0 {
				l += s
			}
		}
		k++
		s <<= 1
	}
	return l
}

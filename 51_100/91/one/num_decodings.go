package one

// 本题最主要的就是要考虑0
func NumDecodings(s string) int {
	// 动态规划
	// last = dp[i-2]
	last := 0
	ans := 1
	pn := uint8(0) // 前一个数字
	for i := range s {
		num := s[i] - '0'
		// 不可能有编码
		if num == 0 && (pn == 0 || pn > 2) {
			return 0
		}
		// 没超过范围则可以结合
		if pn > 0 && num > 0 && pn*10+num <= 26 {
			ans, last = ans+last, ans
		} else if num == 0 {
			// num == 0 那么只能结合pn一起
			ans = last
		} else {
			last = ans
		}
		pn = num
	}
	return ans
}

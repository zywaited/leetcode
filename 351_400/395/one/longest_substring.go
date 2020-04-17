package one

func LongestSubstring(s string, k int) int {
	var cnt func(int, int) int
	cnt = func(si int, ei int) int {
		// 统计数量
		nums := make([]int, 26)
		for i := si; i <= ei; i++ {
			nums[s[i]-'a']++
		}
		// 剔除不能满足的字符
		li := si
		ri := ei
		for ; li <= ei && nums[s[li]-'a'] < k; li++ {
		}
		for ; ri >= li && nums[s[ri]-'a'] < k; ri-- {
		}
		if li > ri {
			return 0
		}
		// 分治
		// 找到分割点
		mi := li
		for ; mi <= ri && nums[s[mi]-'a'] >= k; mi++ {
		}
		// 全都满足
		if mi > ri {
			return ri - li + 1
		}
		return max(cnt(li, mi-1), cnt(mi, ri))
	}
	return cnt(0, len(s)-1)
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}

package one

import "math"

type count map[byte]int

func MinWindow(s string, t string) string {
	sl := len(s)
	tl := len(t)
	if sl < tl || tl == 0 {
		return ""
	}
	left := 0
	right := 0
	minLeft := 0
	ml := math.MaxInt32
	match := 0
	// 以空间换时间
	// 计数器
	window := make(count)
	need := make(count)
	for i := range t {
		// 可能出现重复的字符
		need[t[i]]++
	}
	// t中字符出现顺序
	link := make([]int, sl)
	li := 0
	lk := 0
	for right < sl {
		// 判断是否在t中
		rc := s[right]
		if _, ok := need[rc]; ok {
			window[rc]++
			// 是否匹配
			if window[rc] <= need[rc] {
				match++
			}
			// 顺序记录
			link[lk] = right
			lk++
		}
		right++
		// 已找到
		for match == tl {
			if right-left < ml {
				minLeft = left
				ml = right - left
			}
			// 当前字符是否在t中
			lc := s[left]
			if _, ok := need[lc]; ok {
				window[lc]--
				if window[lc] < need[lc] {
					match--
				}
				// 如果只有1个，则直接退出
				if tl == 1 {
					return string(lc)
				}
				li++
			}
			left = link[li]
		}
	}
	if ml == math.MaxInt32 {
		return ""
	}
	return s[minLeft : minLeft+ml]
}

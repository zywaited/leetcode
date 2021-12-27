package one

func LongestPrefix(s string) string {
	l := len(s)
	next := make([]int, l)
	next[0] = -1
	k := -1
	i := 0
	for i < l-1 {
		if k == -1 || s[i] == s[k] {
			k++
			i++
			next[i] = k
			continue
		}
		k = next[k]
	}
	// next计算的是下一次比较应该跳到的索引
	// 所以这里要确认下数据是否正确
	e := 0
	for i = next[l-1]; i >= 0; i-- {
		if s[i] == s[l-e-1] {
			e++
			continue
		}
		e = 0
	}
	return s[0:e]
}

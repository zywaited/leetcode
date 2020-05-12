package one

func CanConstruct(s string, k int) bool {
	// 不能满足非空
	if k > len(s) {
		return false
	}
	// 每个字符组成回文
	if k == len(s) {
		return true
	}
	max := 0
	// 存在的字符
	exists := 0
	for i := range s {
		j := 1 << (s[i] - 'a')
		if exists&j == 0 {
			exists |= j
			continue
		}
		max++
		exists ^= j
	}
	// 剩余的单个字符数量
	exn := 0
	for exists > 0 {
		exn++
		exists &= exists - 1
	}
	// 单个字符
	if k < exn {
		return false
	}
	return k-exn <= max<<1
}

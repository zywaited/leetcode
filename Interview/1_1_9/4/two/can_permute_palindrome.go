package two

// 题目只说了字母，其他案列上有各种字符
// 那ASCII 128字符
func CanPermutePalindrome(s string) bool {
	hn := uint64(0) // 64-128
	ln := uint64(0) // 0-63
	for i := range s {
		if s[i] >= 64 {
			hn ^= 1 << (s[i] - 64)
		} else {
			ln ^= 1 << s[i]
		}
	}
	// 不用计算所有位1的个数
	if hn > 0 {
		hn &= hn - 1
	}
	if ln > 0 {
		ln &= ln - 1
	}
	return hn+ln <= 1
}

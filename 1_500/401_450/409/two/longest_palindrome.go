package two

// 空间O(1)
func LongestPalindrome(s string) int {
	nums := uint64(0)
	tn := uint64(0) // 临时数据
	ans := 0
	for index := range s {
		if s[index] <= 'Z' {
			tn = 1 << (s[index] - 'A')
		} else {
			tn = 1 << (s[index] - 'a' + 26)
		}
		if tn&nums == 0 {
			nums |= tn
			continue
		}
		ans += 2
		nums ^= tn
	}
	if ans < len(s) {
		ans++
	}
	return ans
}

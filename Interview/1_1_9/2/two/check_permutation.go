package one

// 计算和
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	n1 := 0
	n2 := 0
	for i := range s1 {
		n1 += 1 << (s1[i] - 'a')
		n2 += 1 << (s2[i] - 'a')
	}
	return n1 == n2
}

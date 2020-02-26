package one

// 借助map(长度就是字符，也就是26个)
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	nums := make(map[byte]int)
	for i := range s1 {
		nums[s1[i]]++
		nums[s2[i]]--
	}
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

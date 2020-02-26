package one

func IsUnique(astr string) bool {
	// 位运算
	ex := 0
	for i := range astr {
		n := 1 << uint(astr[i]-'a')
		if ex&n != 0 {
			return false
		}
		ex |= n
	}
	return true
}

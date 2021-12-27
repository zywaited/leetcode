package one

/**
 * 转换为二进制异或处理即可
 */
func FindTheDifference(s string, t string) byte {
	r := byte(0)
	sl := len(s)
	tl := len(t)
	for i := 0; i < tl; i++ {
		if i < sl {
			r ^= s[i]
		}
		r ^= t[i]
	}
	return r
}
